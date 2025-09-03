
import 'package:get_it/get_it.dart';

import 'features/shopping_assistant/data/data_sources/remote/product_remote_data_source.dart';
import 'features/shopping_assistant/data/data_sources/local/product_local_data_source.dart';
import 'features/shopping_assistant/data/repositoriesImpl/repository.dart';
import 'features/shopping_assistant/domain/repositories/prompt_repository.dart';
import 'features/shopping_assistant/domain/usecases/send_prompt.dart';
import 'features/shopping_assistant/presentation/bloc/chat_bloc.dart';

final GetIt sl = GetIt.instance;

Future<void> init() async {
  // Features - Shopping Assistant
  // Bloc
  sl.registerFactory(() => ChatBloc(sendPrompt: sl()));

  // Use cases
  sl.registerLazySingleton(() => SendPrompt(sl()));
  sl.registerLazySingleton<ProductRepository>(
    () => ProductRepositoryImpl(
      remoteDataSource: sl(),
      localDataSource: sl(),
    ),
  );
  // Data sources
  sl.registerLazySingleton<ProductRemoteDataSource>(
    () => ProductRemoteDataSourceImpl(),
  );
  sl.registerLazySingleton<ProductLocalDataSource>(
    () => ProductLocalDataSourceImpl(),
  );

  // External
  // sl.registerLazySingleton(() => http.Client());
}