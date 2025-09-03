class LocalizationStore {
  Map<String, String> _active = {};
  void update(Map<String, String> dict) => _active = dict;
  String getText(String key) => _active[key] ?? key;
}

final localizationStore = LocalizationStore();

String getText(String key) => localizationStore.getText(key);
