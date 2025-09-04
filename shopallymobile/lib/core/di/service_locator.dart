import 'package:get_it/get_it.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'package:shopallymobile/auth_feature/data/datasource/localdatasource.dart';
import 'package:shopallymobile/auth_feature/data/datasource/remotedatasource.dart';
import 'package:shopallymobile/auth_feature/data/repositoryImpl/repositoryImpl.dart';
import 'package:shopallymobile/auth_feature/domain/repositories/user_repo.dart';
import 'package:shopallymobile/core/localization/language_bloc.dart';
import 'package:shopallymobile/core/localization/translation_loader.dart';

final sl = GetIt.instance;

Future<void> initDependencies() async {
  // Async singletons
  final prefs = await SharedPreferences.getInstance();
  sl.registerSingleton<SharedPreferences>(prefs);

  // Core services
  sl.registerLazySingleton<TranslationLoader>(() => TranslationLoader());

  // Data sources
  sl.registerLazySingleton<RemoteDataSource>(() => RemoteDataSourceImpl());
  sl.registerLazySingleton<LocalDataSource>(() => LocalDataSourceImpl(sl()));

  // Repositories (bind to abstractions)
  sl.registerLazySingleton<UserRepository>(() => UserAuthRepositoryImpl(
        remoteDataSource: sl<RemoteDataSource>(),
        localDataSource: sl<LocalDataSource>(),
      ));

  // Blocs
  sl.registerFactory<LanguageBloc>(() => LanguageBloc(sl<TranslationLoader>()));
}
