import 'package:flutter/material.dart';

class AppColors {
  // Backgrounds
  static const Color backgroundTop = Color(0xFFF8FEFB);
  static const Color backgroundBottom = Color(0xFFFFFFFF);
  static const Color cardBackground = Color(0xFFFFFFFF);
  static const Color cardBorder = Color(0xFFE0E0E0);

  // Text
  static const Color textPrimary = Color(0xFF000000);
  static const Color textSecondary = Color(0xFF4A4A4A);
  static const Color textLightGray = Color(0xFF7A7A7A);

  // Branding (Accio logo)
  static const Color accioTeal = Color(0xFF00D4C3);
  static const Color accioGreen = Color(0xFF00E676);
  static const Color accioBlue = Color(0xFF2979FF);

  // Buttons
  static const Color buttonBlack = Color(0xFF000000);
  static const Color buttonText = Color(0xFFFFFFFF);

  // Navigation
  static const Color navInactive = Color(0xFF7A7A7A);
  static const Color navActive = Color(0xFF00C389);

  // Tags / Chips
  static const Color chipBorder = Color(0xFFE0E0E0);
  static const Color chipText = Color(0xFF000000);
  static const Color chipActiveBg = Color(0xFFE6F8F3);
  static const Color chipActiveText = Color(0xFF00C389);

  // Rating
  static const Color starFilled = Color(0xFF000000);

  // Gradient (used in bottom search input glow)
  static const List<Color> inputGlowGradient = [
    Color.fromRGBO(255, 211, 0, 1),
    Color.fromARGB(255, 195, 247, 9),
  ];
}
