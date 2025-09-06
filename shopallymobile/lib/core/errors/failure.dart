import 'package:equatable/equatable.dart';

abstract class Failure extends Equatable {
  final String message;
  
  const Failure(this.message);

  @override
  List<Object?> get props => [message];
}

class ServerFailure extends Failure {
  const ServerFailure({String? message}) : super(message ?? 'Server Failure');
}
class SaveProductFailure extends Failure {
  const SaveProductFailure({String? message}) : super(message ?? 'Save Product Failure');
}
class DatabaseFailure extends Failure {
  const DatabaseFailure({String? message}) : super(message ?? 'Database Failure');
}

class SavedProductError extends Failure {
  const SavedProductError({String? message}) : super(message ?? 'Saved Product Failure');
}

