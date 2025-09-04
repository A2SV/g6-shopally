import 'package:flutter/material.dart';

class PriceAlert extends StatefulWidget {
  const PriceAlert({super.key});

  @override
  State<PriceAlert> createState() => _PriceAlertState();
}

class _PriceAlertState extends State<PriceAlert> {
  bool _isAlertOn = false;

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: 40, // Constrain the width
      height: 24, // Constrain the height
      child: Transform.scale(
        scale: 0.5, // Adjusted scale for better fit
        alignment: Alignment.center,
        child: Switch(
          value: _isAlertOn,
          onChanged: (value) {
            setState(() {
              _isAlertOn = value;
            });
          },
          activeTrackColor: Colors.green,
          activeColor: Colors.white,
        ),
      ),
    );
  }
}