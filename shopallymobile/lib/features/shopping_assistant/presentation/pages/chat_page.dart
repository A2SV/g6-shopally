import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../core/constants/ui_constants.dart';
import '../bloc/chat_bloc.dart';
import '../widgets/prompt_input.dart';
import '../widgets/suggestion_box.dart';


class ChatPage extends StatefulWidget {
  const ChatPage({super.key});

  @override
  State<ChatPage> createState() => _ChatPageState();
}

class _ChatPageState extends State<ChatPage> {
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
    Navigator.of(context).pushReplacementNamed('/detailChat' , arguments: userInput);
    descriptionController.clear();
  }



  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.backgroundTop,
      body: Column(
        children: [
          const Expanded(
            child: Padding(
              padding: EdgeInsets.only(left: 24.0, right: 24.0),
              child: Center(
                child: Text(
                  'Get the product recommendations you need!',
                  textAlign: TextAlign.center,
                  style: TextStyle(
                      color: Colors.black,
                      fontSize: 24,
                      fontWeight: FontWeight.bold),
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
      bottomNavigationBar: BottomNavigationBar(
        backgroundColor: AppColors.backgroundBottom,
        selectedItemColor: Color.fromRGBO(255, 211, 0, 1),
        unselectedItemColor: Colors.black54,
        type: BottomNavigationBarType.fixed,
        currentIndex: _currentIndex,
        onTap: (index) {
          setState(() {
            _currentIndex = index;
          });
          debugPrint('Navigation index tapped: $index');
        },
        items: const [
          BottomNavigationBarItem(icon: Icon(Icons.search ), label: 'Home'),
          BottomNavigationBarItem(icon: Icon(Icons.save_alt_outlined), label: 'Saved'),
          BottomNavigationBarItem(icon: Icon(Icons.compare_arrows_outlined), label: 'Compare'),
          BottomNavigationBarItem(icon: Icon(Icons.person_outline), label: 'Profile'),
        ],
      ),
    );
  }
}
