import 'package:get_it/get_it.dart';
import 'package:internet_connection_checker/internet_connection_checker.dart';

import 'package:shopallymobile/features/comparing/core/network/network_info.dart';
import 'package:shopallymobile/features/comparing/data/datasources/database_shopally.dart';
import 'package:shopallymobile/features/comparing/data/datasources/local_data_source.dart';
import 'package:shopallymobile/features/comparing/data/datasources/remote_data_source.dart';
import 'package:shopallymobile/features/comparing/data/repository/repository_impl.dart';
import 'package:shopallymobile/features/comparing/domain/repository/repository.dart';
import 'package:shopallymobile/features/comparing/domain/usecases/clear_products_useCase.dart';
import 'package:shopallymobile/features/comparing/domain/usecases/compare_products_usecase.dart';
import 'package:shopallymobile/features/comparing/domain/usecases/get_products_for_comparison_useCase.dart';
import 'package:shopallymobile/features/comparing/domain/usecases/remove_from_compare_usecase.dart';
import 'package:shopallymobile/features/comparing/presentation/bloc/compare_bloc.dart';
import 'package:http/http.dart' as http;

// In injection_container.dart
final sl = GetIt.instance;

Future<void> init() async {
  // External dependencies
  sl.registerLazySingleton(() => http.Client());
  sl.registerLazySingleton<InternetConnectionChecker>(
    () => InternetConnectionChecker.createInstance(),
  );

  // Core
  sl.registerLazySingleton<NetworkInfo>(() => NetworkInfoImpl(sl()));

  // Database singleton
  sl.registerLazySingleton<DatabaseShopally>(() => DatabaseShopally.instance);

  // Local data source
  sl.registerLazySingleton<LocalDataSource>(
    () => LocalDataSourceImpl(dbHelper: sl<DatabaseShopally>()),
  );

  // Remote data source
  sl.registerLazySingleton<RemoteDataSource>(
    () => RemoteDataSourceImpl(client: sl()),
  );

  // Repository
  sl.registerLazySingleton<Repository>(
    () => RepositoryImpl(
      localDataSource: sl(),
      remoteDataSource: sl(),
      networkInfo: sl(),
    ),
  );

  // Use cases
  sl.registerLazySingleton(() => CompareProductsUseCase(repository: sl()));
  sl.registerLazySingleton(() => GetProductsForComparison(repository: sl()));
  sl.registerLazySingleton(() => RemoveFromCompareUseCase(repository: sl()));
  sl.registerLazySingleton(() => ClearProductsUseCase(repository: sl()));

  // Bloc
  sl.registerFactory(
    () => CompareBloc(
      compareProductsUseCase: sl(),
      getProductsForComparisonUseCase: sl(),
      removeProductFromComparisonUseCase: sl(),
      clearProductsUseCase: sl(),
    ),
  );
}
