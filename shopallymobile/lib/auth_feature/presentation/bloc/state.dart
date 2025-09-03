import 'package:shopallymobile/auth_feature/domain/entities/user.dart';

abstract class UserAuthState {}

class InitialState extends UserAuthState {}

class LoadingState extends UserAuthState {}

class SuccessState extends UserAuthState {
  final AuthUser? user;
  SuccessState({required this.user});
}

class ErrorState extends UserAuthState {
  final String message;
  ErrorState({required this.message});
}
