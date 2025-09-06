import 'package:shopallymobile/features/auth_feature/domain/entities/user.dart';

abstract class UserRepository {
  Future<AuthUser?> signinWithGoogle();
  Future<void> signout();
  Future<AuthUser?> getCurrentUser();
  Future<AuthUser?> updateLanguage(String language);
  Future<AuthUser?> updateCurrency(String currency);
  Future<AuthUser?> updateNotification(bool toggle);
}
