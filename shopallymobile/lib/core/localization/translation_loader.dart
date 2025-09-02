import 'dart:convert';
import 'package:flutter/services.dart' show rootBundle;

class TranslationLoader {
  Map<String, Map<String, String>> _translations = {};

  Future<void> load() async {
    final raw = await rootBundle.loadString('assets/translation.json');
    final data = json.decode(raw) as Map<String, dynamic>;
    _translations = data.map(
      (lang, map) => MapEntry(
        lang,
        (map as Map<String, dynamic>).map((k, v) => MapEntry(k, v.toString())),
      ),
    );
  }

  Map<String, String> forLang(String code) => _translations[code] ?? {};
  bool get isReady => _translations.isNotEmpty;
}
