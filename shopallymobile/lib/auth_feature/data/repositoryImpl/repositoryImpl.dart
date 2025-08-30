import 'package:shopallymobile/auth_feature/data/datasource/localdatasource.dart';
import 'package:shopallymobile/auth_feature/data/datasource/remotedatasource.dart';
import 'package:shopallymobile/auth_feature/domain/entities/user.dart';
import 'package:shopallymobile/auth_feature/domain/repositories/user_repo.dart';

class UserAuthRepositoryImpl implements UserRepository {
  final RemoteDataSource remoteDataSource;
  final LocalDataSource localDataSource;

  UserAuthRepositoryImpl({
    required this.remoteDataSource,
    required this.localDataSource,
  });

  @override
  Future<AuthUser> signinWithGoogle() async {
    final user = await remoteDataSource.signinWithGoogle();
    await localDataSource.cacheUser(user);
    return user;
  }

  @override
  Future<void> signout() async {
    await remoteDataSource.signout();
    await localDataSource.clearUser();
  }

  @override
  Future<AuthUser?> getCurrentUser() async {
    final user = await localDataSource.getCachedUser();
    return user;
  }

  @override
  Future<AuthUser?> updateLanguage(String language) async {
    final user = await localDataSource.updateLanguage(language);
    return user;
  }

  @override
  Future<AuthUser?> updateCurrency(String currency) async {
    final user = await localDataSource.updateCurrency(currency);
    return user;
  }

  @override
  Future<AuthUser?> updateNotification(bool toggle) async{
    final user = await localDataSource.updateNotification(toggle);
    return user;
  }
}
