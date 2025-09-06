import 'package:get_it/get_it.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'package:shopallymobile/features/auth_feature/data/datasource/localdatasource.dart';
import 'package:shopallymobile/features/auth_feature/data/datasource/remotedatasource.dart';
import 'package:shopallymobile/features/auth_feature/data/repositoryImpl/repositoryImpl.dart';
import 'package:shopallymobile/features/auth_feature/domain/repositories/user_repo.dart';
import 'package:shopallymobile/core/localization/language_bloc.dart';
import 'package:shopallymobile/core/localization/translation_loader.dart';
import 'package:connectivity_plus/connectivity_plus.dart';

// Shopping Assistant imports
import 'package:shopallymobile/features/shopping_assistant/data/data_sources/remote/product_remote_data_source.dart';
import 'package:shopallymobile/features/shopping_assistant/data/data_sources/local/product_local_data_source.dart'
    as sa_local;
import 'package:shopallymobile/features/shopping_assistant/data/repositoriesImpl/repository.dart';
import 'package:shopallymobile/features/shopping_assistant/domain/repositories/prompt_repository.dart';
import 'package:shopallymobile/features/shopping_assistant/domain/usecases/send_prompt.dart';
import 'package:shopallymobile/features/shopping_assistant/presentation/bloc/chat_bloc.dart';

// Saved Items imports (aliased to avoid name clash with SA local DS)
import 'package:shopallymobile/features/saveditem/data/data_sources/local_data/local_data_source.dart'
    as saved_local;
import 'package:shopallymobile/features/saveditem/data/data_sources/local_data/local_data_source_impl.dart'
    as saved_impl;
import 'package:shopallymobile/features/saveditem/data/repositoriesImpl/saved_item_repositories_impl.dart';
import 'package:shopallymobile/features/saveditem/domain/repositories/saved_item_repositories.dart';
import 'package:shopallymobile/features/saveditem/presentation/bloc/bloc/saved_product_bloc.dart';
import 'package:shopallymobile/core/databasehelper/database_helper.dart';
import 'package:shopallymobile/core/network/network_info.dart';

final sl = GetIt.instance;

Future<void> initDependencies() async {
  // Async singletons
  final prefs = await SharedPreferences.getInstance();
  sl.registerSingleton<SharedPreferences>(prefs);

  // Core services
  sl.registerLazySingleton<TranslationLoader>(() => TranslationLoader());

  // External
  sl.registerLazySingleton<Connectivity>(() => Connectivity());
  sl.registerLazySingleton<NetworkInfo>(() => NetworkInfoImpl(sl()));
  sl.registerLazySingleton<DatabaseHelper>(() => DatabaseHelper());

  // Data sources
  sl.registerLazySingleton<RemoteDataSource>(() => RemoteDataSourceImpl());
  sl.registerLazySingleton<LocalDataSource>(() => LocalDataSourceImpl(sl()));

  // Shopping Assistant data sources
  sl.registerLazySingleton<ProductRemoteDataSource>(
      () => ProductRemoteDataSourceImpl());
  sl.registerLazySingleton<sa_local.ProductLocalDataSource>(
      () => sa_local.ProductLocalDataSourceImpl());

  // Repositories (bind to abstractions)
  sl.registerLazySingleton<UserRepository>(() => UserAuthRepositoryImpl(
        remoteDataSource: sl<RemoteDataSource>(),
        localDataSource: sl<LocalDataSource>(),
      ));

  // Shopping Assistant repository
  sl.registerLazySingleton<ProductRepository>(() => ProductRepositoryImpl(
        remoteDataSource: sl<ProductRemoteDataSource>(),
        localDataSource: sl<sa_local.ProductLocalDataSource>(),
      ));

  // Saved Items repository
  sl.registerLazySingleton<saved_local.ProductLocalDataSource>(
      () => saved_impl.ProductLocalDataSourceImpl());
  sl.registerLazySingleton<SavedItemsRepository>(
      () => SavedItemsRepositoryImpl(localDataSource: sl()));

  // Blocs
  sl.registerFactory<LanguageBloc>(() => LanguageBloc(sl<TranslationLoader>()));
  sl.registerLazySingleton<SendPrompt>(() => SendPrompt(sl<ProductRepository>()));
  sl.registerFactory<ChatBloc>(() => ChatBloc(sendPrompt: sl<SendPrompt>()));
  sl.registerFactory<SavedProductBloc>(() => SavedProductBloc(sl<SavedItemsRepository>()));
}
