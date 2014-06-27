package dbrary

import scala.util.control.Exception.catching
import play.api.libs.json
import macros._

/** Enumerations which reflect types in the database.
  * Any PGEnum should exactly match the correspending database type. */
abstract class PGEnum(name : String) extends Enumeration {
  implicit val sqlType : SQLType[Value] =
    SQLType[Value](name, classOf[Value])(withNameOpt, _.toString)
  implicit val jsonFormat : json.Format[Value] = new json.Format[Value] {
    def writes(v : Value) = json.JsNumber(v.id)
    def reads(j : json.JsValue) = j match {
      case json.JsNumber(i) if i.isValidInt && i >= 0 && i < maxId => json.JsSuccess(apply(i.toInt))
      case _ => json.JsError("error.expected.jsnumber")
    }
  }
  final def withNameOpt(name : String) : Option[Value] =
    catching(classOf[NoSuchElementException]).opt(withName(name))
  final def fromString(s : String) : Option[Value] =
    withNameOpt(s) orElse
      Maybe.toInt(s).flatMap(i => catching(classOf[NoSuchElementException]).opt(apply(i)))
}

object PGEnum {
  import scala.language.experimental.macros
  import scala.reflect.macros.blackbox.Context
  import Macro._

  /* This is not very useful as it can only create structural values rather than top-level objects */
  def make(enumName : String) : PGEnum = macro makeImpl

  def makeImpl(c : Context)(enumName : c.Expr[String]) : c.Expr[PGEnum] = {
    import c.universe._
    val name = getString(c)(enumName)
    val labels = Connection.Static.enumLabels(name)
    val obj = TermName(name)

    c.Expr(Block(
      List(ModuleDef(Modifiers(), obj, Template(
        List(Ident(c.mirror.staticClass("dbrary.PGEnum"))),
        noSelfType,
        DefDef(Modifiers(), termNames.CONSTRUCTOR, Nil, List(Nil), TypeTree(), Block(
          List(Apply(Select(Super(This(typeNames.EMPTY), typeNames.EMPTY), termNames.CONSTRUCTOR), 
            List(enumName.tree))),
          Literal(Constant(()))))
        :: labels.map(l => 
          ValDef(Modifiers(), TermName(l), TypeTree(), Select(This(TypeName("")), TermName("Value"))))
      ))),
      Ident(obj)
    ))
  }
}
