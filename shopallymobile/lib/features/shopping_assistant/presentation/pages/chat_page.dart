import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:flutter/material.dart';
import 'package:shopallymobile/features/compare/presentation/pages/compare_page.dart';
import 'package:shopallymobile/features/comparing/presentation/pages/products_for_comparing.dart';
import 'package:shopallymobile/features/saveditem/presentation/pages/savedpage.dart';
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
    const ProductsForComparing(),
  ];

  @override
  Widget build(BuildContext context) {
    final iconColor = Theme.of(context).textTheme.bodyLarge?.color;
    final navItems = <Widget>[
      Icon(Icons.search, size: 30, color: iconColor),
      Icon(Icons.favorite_border, size: 30, color: iconColor),
      Icon(Icons.swap_horiz, size: 30, color: iconColor),
      Icon(Icons.person_outline, size: 30, color: iconColor),
    ];
    return Scaffold(
      backgroundColor: Theme.of(context).scaffoldBackgroundColor,
      body: _pages[_currentIndex],
      bottomNavigationBar: CurvedNavigationBar(
        backgroundColor: Colors.transparent,
        color: Theme.of(context).cardColor,
        buttonBackgroundColor: Theme.of(context).cardColor,
        height: 60.0,
        items: navItems,
        onTap: (index) {
          if (index == 3) {
            // Navigate to the app-level Profile route so it receives theme props
            Navigator.pushNamed(context, '/profile');
            return;
          }
          setState(() {
            _currentIndex = index;
          });
        },
      ),
    );
  }
}
