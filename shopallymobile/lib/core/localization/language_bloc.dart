import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'language_event.dart';
import 'language_state.dart';
import 'translation_loader.dart';

class LanguageBloc extends Bloc<LanguageEvent, LanguageState> {
  final TranslationLoader loader;
  LanguageBloc(this.loader) : super(LanguageInitial()) {
    on<LoadLanguageEvent>(_onLoad);
    on<SetLanguageEvent>(_onSet);
  }

  Future<void> _onLoad(
    LoadLanguageEvent event,
    Emitter<LanguageState> emit,
  ) async {
    emit(LanguageLoading());
    try {
      await loader.load();
      final prefs = await SharedPreferences.getInstance();
      final saved = event.initialCode ?? prefs.getString('lang_code') ?? 'en';
      emit(LanguageLoaded(code: saved, dict: loader.forLang(saved)));
    } catch (e) {
      emit(LanguageError(e.toString()));
    }
  }

  Future<void> _onSet(
    SetLanguageEvent event,
    Emitter<LanguageState> emit,
  ) async {
    try {
      final prefs = await SharedPreferences.getInstance();
      await prefs.setString('lang_code', event.code);
      emit(LanguageLoaded(code: event.code, dict: loader.forLang(event.code)));
    } catch (e) {
      emit(LanguageError(e.toString()));
    }
  }
}
