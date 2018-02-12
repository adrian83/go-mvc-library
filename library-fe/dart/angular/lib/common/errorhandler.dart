import 'errors.dart';

import 'package:logging/logging.dart';

class ErrorHandler {
  static final Logger LOGGER = new Logger('ErrorHandler');

  List<ValidationError> _validationErrors;
  ServerError _serverError;

  void handleError(e) {
    if (e is ValidationErrors) {
      this.validationErrors = e.validationErrors;
      LOGGER.info("Validation errors: $validationErrors");
      return;
    } else if (e is ServerError) {
      this._serverError = e;
      LOGGER.info("Server error: $_serverError");
      return;
    }

    LOGGER.info("Errors: $e");
  }

  void set validationErrors(List<ValidationError> errors){
    this._validationErrors = errors;
  }

  List<ValidationError> get validationErrors => this._validationErrors == null
      ? new List<ValidationError>()
      : this._validationErrors;

  ServerError get serverError => _serverError;

  void cleanValidationErrors(){
    this._validationErrors = new List<ValidationError>();
  }

}
