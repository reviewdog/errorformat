// scalac -Ywarn-unused scalac.scala &> ../../scalac.in
// sbt -no-colors compile > ../../sbt.in
object F {
  private val unused = 1
  private def unusedF = {}
  error
}
