import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shopallymobile/features/auth_feature/domain/repositories/user_repo.dart';
import 'package:shopallymobile/features/auth_feature/presentation/pages/profilepage.dart';
import 'package:shopallymobile/core/di/service_locator.dart';
import 'package:shopallymobile/core/localization/language_bloc.dart';
import 'package:shopallymobile/core/localization/language_event.dart';
import 'package:shopallymobile/core/localization/language_state.dart';
import 'package:shopallymobile/core/localization/localization_store.dart';
import 'package:shopallymobile/core/theme/theme_data.dart';
import 'package:shopallymobile/features/compare/presentation/pages/compare_page.dart';
import 'package:shopallymobile/features/saveditem/presentation/pages/savedpage.dart';
import 'package:shopallymobile/features/shopping_assistant/presentation/bloc/chat_bloc.dart';
import 'package:shopallymobile/features/saveditem/presentation/bloc/bloc/saved_product_bloc.dart';
import 'package:shopallymobile/features/shopping_assistant/presentation/pages/chat_page.dart';
import 'package:shopallymobile/features/shopping_assistant/presentation/pages/chat_response_page.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:shopallymobile/features/comparing/presentation/bloc/compare_bloc.dart';
import 'package:shopallymobile/features/comparing/presentation/pages/products_for_comparing.dart';
import 'package:shopallymobile/features/comparing/injection_container.dart' as compare_di;

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();

  // Initialize service locator
  await initDependencies();
  // Initialize comparing feature DI
  await compare_di.init();

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
          BlocProvider<ChatBloc>(
            create: (_) => sl<ChatBloc>(),
          ),
          BlocProvider<SavedProductBloc>(
            create: (_) => sl<SavedProductBloc>(),
          ),
          BlocProvider<CompareBloc>(
            create: (_) => compare_di.sl<CompareBloc>(),
          ),
        ],
        child: const ShopAllyApp(),
      ),
    ),
  );
}

class ShopAllyApp extends StatefulWidget {
  const ShopAllyApp({super.key});

  @override
  State<ShopAllyApp> createState() => _MyAppState();
}

class _MyAppState extends State<ShopAllyApp> {
  ThemeMode _themeMode = ThemeMode.light; // default: dark mode OFF

  @override
  void initState() {
    super.initState();
    _loadSavedTheme();
  }

  Future<void> _loadSavedTheme() async {
    final prefs = await SharedPreferences.getInstance();
    final isDark = prefs.getBool('dark_mode') ?? false;
    setState(() {
      _themeMode = isDark ? ThemeMode.dark : ThemeMode.light;
    });
  }

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

          routes:{
            '/chat':(_)=> ChatPage(),
            '/detailChat': (context) {
          final prompt = ModalRoute.of(context)?.settings.arguments as String? ?? '';
          return ChatResponsePage(prompt: prompt);
        },
        '/saved': (_) => const Savedpage(),
        '/compare': (_) => const ComparePage(),
        '/compare-products': (_) => const ProductsForComparing(),
        '/profile': (_) =>  ProfilePage(
              userRepository: sl<UserRepository>(),
              isDark: _themeMode == ThemeMode.dark,
              onDarkChanged: (val) {
                setState(() {
                  _themeMode = val ? ThemeMode.dark : ThemeMode.light;
                });
                // Persist theme preference
                SharedPreferences.getInstance().then(
                  (p) => p.setBool('dark_mode', val),
                );
              },
            ),
          },
          home: BlocBuilder<ChatBloc, ChatState>(
            builder: (context, state) {
              if (state is ChatLoadingState) {
                return const Center(child: CircularProgressIndicator());
              } else if (state is ChatErrorState) {

                return Scaffold(body: Center(child: Text(state.message)));
              } else {
                return const ChatPage();
              }
            },
          ),
        );
      },
    );
  }
}
