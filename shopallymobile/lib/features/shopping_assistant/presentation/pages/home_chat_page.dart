import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shopallymobile/core/localization/localization_store.dart';
import '../bloc/chat_bloc.dart';
import '../widgets/prompt_input.dart';
import '../widgets/suggestion_box.dart';

class HomeChatPage extends StatefulWidget {
  const HomeChatPage({super.key});

  @override
  State<HomeChatPage> createState() => _HomeChatPageState();
}

class _HomeChatPageState extends State<HomeChatPage> {
  final TextEditingController descriptionController = TextEditingController();
  int _currentIndex = 0;
  bool showSuggestions = true;

  @override
  void initState() {
    super.initState();
    descriptionController.addListener(() {
      setState(() {
        showSuggestions = descriptionController.text.isEmpty;
      });
    });
  }

  // @override
  // void dispose() {
  //   descriptionController.dispose();
  //   super.dispose();
  // }

  void _handleSubmit() {
    final userInput = descriptionController.text.trim();
    if (userInput.isEmpty) return;
    debugPrint('User submitted: $userInput');
    context.read<ChatBloc>().add(SendPromptRequested(userInput));
    Navigator.of(
      context,
    ).pushReplacementNamed('/detailChat', arguments: userInput);
    descriptionController.clear();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Theme.of(context).scaffoldBackgroundColor,
      body: Column(
        children: [
           Expanded(
            child: Padding(
              padding: EdgeInsets.only(left: 24.0, right: 24.0),
              child: Center(
                child: Text(
                  getText('get_the_product_recommendations_you_need!'),
                  textAlign: TextAlign.center,
                  style: TextStyle(
                    color: Theme.of(context).textTheme.bodyLarge?.color,
                    fontSize: 24,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ),
            ),
          ),
          if (showSuggestions)
            SuggestionBox(descriptionController: descriptionController),
          PromptInput(
            onSubmit: _handleSubmit,
            descriptionController: descriptionController,
          ),
          const SizedBox(height: 10),
        ],
      ),
    );
  }
}
