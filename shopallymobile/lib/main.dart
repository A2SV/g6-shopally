import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'features/shopping_assistant/presentation/bloc/chat_bloc.dart';
import 'features/shopping_assistant/presentation/pages/chat_page.dart';
import 'features/shopping_assistant/presentation/pages/chat_response_page.dart';
import 'injection_container.dart' as di;

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await di.init();

  final ChatBloc chatBloc = di.sl<ChatBloc>();

  runApp(
    BlocProvider<ChatBloc>.value(
      value: chatBloc,
      child: const ShopAllyApp(),
    ),
  );
}

class ShopAllyApp extends StatelessWidget {
  const ShopAllyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'ShopAlly',
      routes: {
        '/chat': (_) => const ChatPage(),
        '/detailChat': (context) {
          final prompt = ModalRoute.of(context)?.settings.arguments as String? ?? '';
          return ChatResponsePage(prompt: prompt);
        },
        // Add more routes here as needed
        // '/error': (_) => ErrorPage(...),
        // '/splash': (_) => SplashScreen(),
      },
      home: BlocBuilder<ChatBloc, ChatState>(
        builder: (context, state) {
            if (state is ChatLoadingState) {
              return const Center(child: CircularProgressIndicator());
            } else if (state is ChatErrorState) {
              return Scaffold(
                body: Center(child: Text(state.message)),
              );
            } else {
              return const ChatPage();
            }
        },
      ),
    );
  }
}