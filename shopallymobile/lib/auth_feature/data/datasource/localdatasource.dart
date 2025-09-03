import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';

import 'package:shopallymobile/auth_feature/data/model/user_model.dart';

abstract class LocalDataSource {
  Future<UserModel?> getCachedUser();
  Future<void> cacheUser(UserModel user);
  Future<void> clearUser();
  Future<UserModel?> updateLanguage(String language);
  Future<UserModel?> updateNotification(bool toggle);
  Future<UserModel?> updateCurrency(String currency);
}

class LocalDataSourceImpl implements LocalDataSource {
  final SharedPreferences prefs;

  LocalDataSourceImpl(this.prefs);

  static const String _cachedUserKey = "CACHED_USER";

  @override
  Future<void> cacheUser(UserModel user) async {
    // Ensure defaults are present when caching
    final model = UserModel(
      email: user.email,
      name: user.name,
      photourl: user.photourl,
      language: user.language ?? 'English',
      currency: user.currency ?? 'USD',
      notifications: user.notifications ?? true,
    );
    final jsonString = jsonEncode(model.tojson());
    await prefs.setString(_cachedUserKey, jsonString);
  }

  @override
  Future<void> clearUser() async {
    await prefs.remove(_cachedUserKey);
  }

  @override
  Future<UserModel?> getCachedUser() async {
    final jsonString = prefs.getString(_cachedUserKey);
    if (jsonString == null) return null;
    final jsonMap = jsonDecode(jsonString);
    print("Helllooooo");
    print(jsonMap);
    return UserModel.fromjson(jsonMap);
  }

  @override
  Future<UserModel?> updateCurrency(String currency) async {
    final jsonString = prefs.getString(_cachedUserKey);
    if (jsonString != null) {
      try {
        final map = jsonDecode(jsonString) as Map<String, dynamic>;
        map['currency'] = currency;
        await prefs.setString(_cachedUserKey, jsonEncode(map));
        return UserModel.fromjson(map);
      } catch (e) {
        return null;
      }
    }
    return null;
  }

  @override
  Future<UserModel?> updateLanguage(String language) async {
    final jsonString = prefs.getString(_cachedUserKey);

    if (jsonString != null) {
      try {
        final map = jsonDecode(jsonString) as Map<String, dynamic>;
        map['language'] = language;
        await prefs.setString(_cachedUserKey, jsonEncode(map));
        return UserModel.fromjson(map);
      } catch (e) {
        return null;
      }
    }
    return null;
  }

  @override
  Future<UserModel?> updateNotification(bool toggle) async {
    final jsonString = prefs.getString(_cachedUserKey);
    if (jsonString != null) {
      try {
        final map = jsonDecode(jsonString) as Map<String, dynamic>;
        map['notifications'] = toggle;
        await prefs.setString(_cachedUserKey, jsonEncode(map));
        return UserModel.fromjson(map);
      } catch (e) {
        return null;
      }
    }
    return null;
  }
}
