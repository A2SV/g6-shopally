import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:shopallymobile/auth_feature/data/datasource/localdatasource.dart';
import 'package:shopallymobile/auth_feature/data/datasource/remotedatasource.dart';
import 'package:shopallymobile/auth_feature/data/repositoryImpl/repositoryImpl.dart';
import 'package:shopallymobile/auth_feature/domain/repositories/user_repo.dart';
import 'package:shopallymobile/auth_feature/presentation/pages/profilepage.dart';

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
      child: const MyApp(),
    ),
  );
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Flutter Demo',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
      ),
      home: ProfilePage(userRepository: context.read<UserRepository>()),
    );
  }
}
