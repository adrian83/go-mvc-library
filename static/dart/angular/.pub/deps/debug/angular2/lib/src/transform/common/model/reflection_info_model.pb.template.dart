// @ignoreProblemForFile UNUSED_IMPORT
// @ignoreProblemForFile UNUSED_SHOWN_NAME
library angular2.src.transform.common.model.proto_reflection_info_model.template.dart;

import 'reflection_info_model.pb.dart';
import 'package:angular2/src/core/reflection/reflection.dart' as _ngRef;
import 'package:protobuf/protobuf.dart';
import 'annotation_model.pb.dart';
import 'parameter_model.pb.dart';
import 'annotation_model.pb.template.dart' as i0;
import 'parameter_model.pb.template.dart' as i1;
export 'reflection_info_model.pb.dart';

var _visited = false;
void initReflector() {
if (_visited) return; _visited = true;
i0.initReflector();
i1.initReflector();
}