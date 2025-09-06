import 'package:connectivity_plus/connectivity_plus.dart';
import 'package:get_it/get_it.dart';
import 'package:http/http.dart' as http;
import 'package:shopallymobile/core/databasehelper/database_helper.dart';
import 'package:shopallymobile/core/network/network_info.dart';
import 'package:shopallymobile/features/saveditem/data/data_sources/local_data/local_data_source.dart';
import 'package:shopallymobile/features/saveditem/data/data_sources/local_data/local_data_source_impl.dart';
import 'package:shopallymobile/features/saveditem/data/repositoriesImpl/saved_item_repositories_impl.dart';
import 'package:shopallymobile/features/saveditem/domain/repositories/saved_item_repositories.dart';
import 'package:shopallymobile/features/saveditem/domain/usecases/get_saved_item.dart';
import 'package:shopallymobile/features/saveditem/domain/usecases/remove_product.dart';
import 'package:shopallymobile/features/saveditem/domain/usecases/save_product.dart';
import 'package:shopallymobile/features/saveditem/presentation/bloc/bloc/saved_product_bloc.dart';

final sl = GetIt.instance;

Future<void> init() async {
  // Features - Saved Product
  // Bloc
  sl.registerFactory(() => SavedProductBloc(sl()));

  // Use cases
  sl.registerLazySingleton(() => GetSavedItems(sl()));
  sl.registerLazySingleton(() => SaveProduct(sl()));
  sl.registerLazySingleton(() => RemoveProduct(sl()));

  // Repository
 sl.registerLazySingleton<SavedItemsRepository>(
  () => SavedItemsRepositoryImpl(localDataSource: sl()),
);

  // Data sources
  sl.registerLazySingleton<ProductLocalDataSource>(
    () => ProductLocalDataSourceImpl(),
  );

  // Core
  sl.registerLazySingleton<NetworkInfo>(() => NetworkInfoImpl(sl()));
  sl.registerLazySingleton(() => DatabaseHelper());

  // External
  sl.registerLazySingleton(() => http.Client());
  sl.registerLazySingleton(() => Connectivity());
}
