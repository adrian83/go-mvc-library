// @ignoreProblemForFile UNUSED_IMPORT
// @ignoreProblemForFile UNUSED_SHOWN_NAME
library angular2.src.transform.common.zone.template.dart;

import 'zone.dart';
import 'package:angular2/src/core/reflection/reflection.dart' as _ngRef;
import 'dart:async';
import 'package:analyzer/analyzer.dart';
import 'package:angular2/src/compiler/offline_compiler.dart';
import 'package:barback/barback.dart';
import 'package:source_span/source_span.dart';
import 'package:angular2/src/compiler/offline_compiler.template.dart' as i0;
export 'zone.dart';

var _visited = false;
void initReflector() {
if (_visited) return; _visited = true;
i0.initReflector();
}
