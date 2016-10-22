// scalac -deprecation -feature -unchecked -Xlint -Ywarn-dead-code -Ywarn-numeric-widen -Ywarn-unused -Ywarn-unused-import -Ywarn-value-discard scalac.scala &> scalac.in
object F {
  private val unused = 1
  private def unusedF = {}
  error
}
