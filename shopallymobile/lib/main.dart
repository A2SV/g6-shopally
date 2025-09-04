import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shopallymobile/auth_feature/domain/repositories/user_repo.dart';
import 'package:shopallymobile/auth_feature/presentation/pages/profilepage.dart';
import 'package:shopallymobile/core/di/service_locator.dart';
import 'package:shopallymobile/core/localization/language_bloc.dart';
import 'package:shopallymobile/core/localization/language_event.dart';
import 'package:shopallymobile/core/localization/language_state.dart';
import 'package:shopallymobile/core/localization/localization_store.dart';
import 'package:shopallymobile/core/theme/theme_data.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();

  // Initialize service locator
  await initDependencies();

  runApp(
    MultiRepositoryProvider(
      providers: [
        // Provide repository from service locator to keep context.read working
        RepositoryProvider<UserRepository>.value(value: sl<UserRepository>()),
      ],
      child: MultiBlocProvider(
        providers: [
          BlocProvider<LanguageBloc>(
            create: (_) =>
                sl<LanguageBloc>()..add(LoadLanguageEvent(initialCode: null)),
          ),
      ],
      child: const MyApp(),
      ),
    ),
  );
}

class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  State<MyApp> createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  ThemeMode _themeMode = ThemeMode.light; // default: dark mode OFF

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<LanguageBloc, LanguageState>(
      builder: (context, state) {
        if (state is LanguageLoaded) {
          localizationStore.update(state.dict);
        }
        return MaterialApp(
          debugShowCheckedModeBanner: false,
          theme: AppThemes.lightTheme,
          darkTheme: AppThemes.darkTheme,
          themeMode: _themeMode,
          home: ProfilePage(
            userRepository: sl<UserRepository>(),
            isDark: _themeMode == ThemeMode.dark,
            onDarkChanged: (val) {
              setState(() {
                _themeMode = val ? ThemeMode.dark : ThemeMode.light;
              });
            },
          ),
        );
      },
    );
  }
}
