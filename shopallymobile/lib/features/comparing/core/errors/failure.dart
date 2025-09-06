// lib/core/error/failures.dart

import 'package:equatable/equatable.dart';

abstract class Failure {
  final String message;

  Failure(this.message);

  @override
  String toString() => message;
}

// Example concrete failures
class ServerFailure extends Failure with EquatableMixin {
  ServerFailure([super.message = 'Server Failure']);

  @override
  List<Object?> get props => [message];
}

class CacheFailure extends Failure with EquatableMixin {
  CacheFailure([super.message = 'Cache Failure']);

  @override
  List<Object?> get props => [message];
}

class NotFoundFailure extends Failure with EquatableMixin {
  NotFoundFailure([super.message = 'Item not found']);

  @override
  List<Object?> get props => [message];
}

class NetworkFailure extends Failure with EquatableMixin {
  NetworkFailure([super.message = 'No internet connection. Please check your network.']);

  @override
  List<Object?> get props => [message];
}