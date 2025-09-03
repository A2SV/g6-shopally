abstract class UserAuthEvent {}

class SignInEvent extends UserAuthEvent {}

class SignOutEvent extends UserAuthEvent {}

class UpdateLanguageEvent extends UserAuthEvent {
  final String language;
  UpdateLanguageEvent(this.language);
}

class GetAuthenticatedUserEvent extends UserAuthEvent {}

class UpdateCurrencyEvent extends UserAuthEvent {
  final String currency;
  UpdateCurrencyEvent(this.currency);
}

class UpdateNotificationEvent extends UserAuthEvent {
  final bool toggle;
  UpdateNotificationEvent(this.toggle);
}
