// @ignoreProblemForFile UNUSED_IMPORT
// @ignoreProblemForFile UNUSED_SHOWN_NAME
library angular2.src.compiler.output.path_util.template.dart;

import 'path_util.dart';
import 'package:angular2/src/core/reflection/reflection.dart' as _ngRef;
import 'package:angular2/src/facade/exceptions.dart' show BaseException;
import 'package:angular2/src/facade/lang.dart' show isPresent, isBlank, RegExpWrapper, Math;
import 'package:angular2/src/facade/exceptions.template.dart' as i0;
import 'package:angular2/src/facade/lang.template.dart' as i1;
export 'path_util.dart';

var _visited = false;
void initReflector() {
if (_visited) return; _visited = true;
i0.initReflector();
i1.initReflector();
}
