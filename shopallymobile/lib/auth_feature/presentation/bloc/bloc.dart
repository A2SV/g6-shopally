import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shopallymobile/auth_feature/domain/repositories/user_repo.dart';
import 'package:shopallymobile/auth_feature/presentation/bloc/event.dart';
import 'package:shopallymobile/auth_feature/presentation/bloc/state.dart';

class UserAuthBloc extends Bloc<UserAuthEvent, UserAuthState> {
  final UserRepository userRepositories;
  UserAuthBloc(this.userRepositories) : super(InitialState()) {
    on<SignInEvent>((event, emit) async {
      emit(InitialState());
      try {
        final user = await userRepositories.signinWithGoogle();
        emit(SuccessState(user: user));
      } catch (e) {
        emit(ErrorState(message: e.toString()));
      }
    });
    on<SignOutEvent>((event, emit) async {
      emit(InitialState());
      try {
        await userRepositories.signout();
        emit(SuccessState(user: null));
      } catch (e) {
        emit(ErrorState(message: e.toString()));
      }
    });

    on<GetAuthenticatedUserEvent>((event, emit) async {
      emit(InitialState());
      try {
        final user = await userRepositories.getCurrentUser();
        emit(SuccessState(user: user));
      } catch (e) {
        emit(ErrorState(message: e.toString()));
      }
    });

     on<UpdateLanguageEvent>((event, emit) async {
      emit(InitialState());
      try{
        final user = await userRepositories.updateLanguage(event.language);
        emit(SuccessState(user: user));
      }
      catch(e){
        emit(ErrorState(message: e.toString()));
      }
    });
    on<UpdateCurrencyEvent>((event, emit) async {
      emit(InitialState());
      try{
        final user = await userRepositories.updateCurrency(event.currency);
        emit(SuccessState(user: user));
      }
      catch(e){
        emit(ErrorState(message: e.toString()));
      }
    });
    on<UpdateNotificationEvent>((event, emit) async {
      emit(InitialState());
      try{
        final user = await userRepositories.updateNotification(event.toggle);
        emit(SuccessState(user: user));
      }
      catch(e){
        emit(ErrorState(message: e.toString()));
      }
    });
  }
}
