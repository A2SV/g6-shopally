import 'package:shopallymobile/auth_feature/domain/entities/user.dart';

class UserModel extends AuthUser {
  UserModel({
    required String name,
    required String email,
    String? photourl,
    String? language,
    String? currency,
    bool? notifications,
  }) : super(
         name: name,
         email: email,
         photourl: photourl,
         language: language,
         currency: currency,
         notifications: notifications,
       );

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
