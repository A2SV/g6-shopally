import 'package:shopallymobile/features/auth_feature/domain/entities/user.dart';

class UserModel extends AuthUser {
  UserModel({
    required super.name,
    required super.email,
    super.photourl,
    super.language,
    super.currency,
    super.notifications,
  });

  factory UserModel.fromjson(Map<String, dynamic> json) {
    return UserModel(
      name: json['name'] as String,
      email: json['email'] as String,
      photourl: json['photourl'] as String?,
      language: (json['language'] as String?) ?? 'English',
      currency: (json['currency'] as String?) ?? 'USD',
      notifications: (json['notifications'] as bool?) ?? true,
    );
  }

  Map<String, dynamic> tojson() {
    return {
      'name': name,
      'email': email,
      'photourl': photourl,
      'language': language ?? 'English',
      'currency': currency ?? 'USD',
      'notifications': notifications ?? true,
    };
  }
}
