abstract class LanguageState {}

class LanguageInitial extends LanguageState {}

class LanguageLoading extends LanguageState {}

class LanguageLoaded extends LanguageState {
  final String code; // e.g., 'en' or 'am'
  final Map<String, String> dict;
  LanguageLoaded({required this.code, required this.dict});
}

class LanguageError extends LanguageState {
  final String message;
  LanguageError(this.message);
}
