// @ignoreProblemForFile UNUSED_IMPORT
// @ignoreProblemForFile UNUSED_SHOWN_NAME
library angular2.instrumentation.template.dart;

import 'instrumentation.dart';
import 'package:angular2/src/core/reflection/reflection.dart' as _ngRef;
import 'src/core/profile/profile.template.dart' as i0;
export 'instrumentation.dart';
export 'src/core/profile/profile.dart' show wtfCreateScope, wtfLeave, wtfStartTimeRange, wtfEndTimeRange, WtfScopeFn;

var _visited = false;
void initReflector() {
if (_visited) return; _visited = true;
i0.initReflector();
}
