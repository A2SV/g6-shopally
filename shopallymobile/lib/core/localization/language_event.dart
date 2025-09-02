abstract class LanguageEvent {}

class LoadLanguageEvent extends LanguageEvent {
  final String? initialCode; // e.g., 'en' or 'am'
  LoadLanguageEvent({this.initialCode});
}

class SetLanguageEvent extends LanguageEvent {
  final String code; // 'en' or 'am'
  SetLanguageEvent(this.code);
}
