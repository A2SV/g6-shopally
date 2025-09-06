import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:flutter/material.dart';
import 'package:shopallymobile/features/compare/presentation/pages/compare_page.dart';
import 'package:shopallymobile/features/profile/presentation/pages/profile_page.dart';
import 'package:shopallymobile/features/saveditem/presentation/pages/savedpage.dart';

import '../../../../core/constants/ui_constants.dart';
import './home_chat_page.dart';

class ChatPage extends StatefulWidget {
  const ChatPage({super.key});

  @override
  State<ChatPage> createState() => _ChatPageState();
}

class _ChatPageState extends State<ChatPage> {
  int _currentIndex = 0;

  final List<Widget> _pages = [
    const HomeChatPage(), 
    const Savedpage(), 
    const ComparePage(), 
    const ProfilePage(), 
  ];

  final _icons = <Widget>[
    const Icon(Icons.search, size: 30),
    const Icon(Icons.favorite_border, size: 30),
    const Icon(Icons.swap_horiz, size: 30),
    const Icon(Icons.person_outline, size: 30),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.backgroundTop,
      body: _pages[_currentIndex],
      bottomNavigationBar: CurvedNavigationBar(
        backgroundColor: Colors.transparent,
        color: const Color.fromARGB(255, 238, 230, 230),
        buttonBackgroundColor: Colors.white,
        height: 60.0,
        items: _icons,
        onTap: (index) {
          setState(() {
            _currentIndex = index;
          });
        },
      ),
    );
  }
}
