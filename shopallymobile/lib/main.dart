import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:shopallymobile/auth_feature/data/datasource/localdatasource.dart';
import 'package:shopallymobile/auth_feature/data/datasource/remotedatasource.dart';
import 'package:shopallymobile/auth_feature/data/repositoryImpl/repositoryImpl.dart';
import 'package:shopallymobile/auth_feature/domain/repositories/user_repo.dart';
import 'package:shopallymobile/auth_feature/presentation/pages/profilepage.dart';
import 'package:shopallymobile/core/localization/language_bloc.dart';
import 'package:shopallymobile/core/localization/language_event.dart';
import 'package:shopallymobile/core/localization/language_state.dart';
import 'package:shopallymobile/core/localization/localization_store.dart';
import 'package:shopallymobile/core/localization/translation_loader.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();

  final prefs = await SharedPreferences.getInstance();

  final userFeatureRepository = UserAuthRepositoryImpl(
    remoteDataSource: RemoteDataSourceImpl(),
    localDataSource: LocalDataSourceImpl(prefs),
  );
  runApp(
    MultiRepositoryProvider(
      providers: [
        RepositoryProvider<UserRepository>.value(value: userFeatureRepository),
      ],
      child: MultiBlocProvider(
        providers: [
          BlocProvider<LanguageBloc>(
            create: (_) =>
                LanguageBloc(TranslationLoader())
                  ..add(LoadLanguageEvent(initialCode: null)),
          ),
        ],
        child: const MyApp(),
      ),
    ),
  );
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<LanguageBloc, LanguageState>(
      builder: (context, state) {
        if (state is LanguageLoaded) {
          localizationStore.update(state.dict);
        }
        return MaterialApp(
          debugShowCheckedModeBanner: false,
          home: ProfilePage(userRepository: context.read<UserRepository>()),
        );
      },
    );
  }
}
