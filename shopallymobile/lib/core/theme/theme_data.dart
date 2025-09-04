import 'package:flutter/material.dart';
import 'colors/light_colors.dart';
import 'colors/dark_colors.dart';

class AppThemes {
  static final lightTheme = ThemeData(
    brightness: Brightness.light,
    scaffoldBackgroundColor: LightColors.background,
    textTheme: const TextTheme(
      bodyLarge: TextStyle(color: LightColors.text),
      bodyMedium: TextStyle(color: LightColors.text),
      bodySmall: TextStyle(color: LightColors.text),
    ),
  );

  static final darkTheme = ThemeData(
    brightness: Brightness.dark,
    scaffoldBackgroundColor: DarkColors.background,
    textTheme: const TextTheme(
      bodyLarge: TextStyle(color: DarkColors.text),
      bodyMedium: TextStyle(color: DarkColors.text),
      bodySmall: TextStyle(color: DarkColors.text),
    ),
  );
}
