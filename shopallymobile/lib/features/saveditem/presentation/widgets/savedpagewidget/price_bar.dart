import 'package:flutter/material.dart';
import 'package:shopallymobile/core/constants/const_color.dart';

class PriceBar extends StatefulWidget {
  const PriceBar({super.key});

  @override
  State<PriceBar> createState() => _PriceBarState();
}

class _PriceBarState extends State<PriceBar> {
  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.all(16.0),
      margin: const EdgeInsets.only(left:  16.0 , right: 16),
      decoration: BoxDecoration(
        color: grey.withOpacity(0.1),
        border: Border(top: BorderSide(color: Colors.grey.shade300)),
        borderRadius: BorderRadius.circular(8.0),
        boxShadow: [
          BoxShadow(
            color: grey.withOpacity(0.1),
            blurRadius: 10.0,
            spreadRadius: 2.0,
            offset: const Offset(0, -2),
          ),
        ],
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: const [
                Text(
                  'Price',
                  style: TextStyle(fontSize: 14, color: Colors.grey),
                ),
                Text(
                  '\$4.50 - \$5.29',
                  style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
                ),
                Text(
                  '100 - 999 pieces:',
                  style: TextStyle(fontSize: 14, color: Colors.grey),
                ),
              ],
            ),
          ),
       
        ],
      ),
    );
  }
}
