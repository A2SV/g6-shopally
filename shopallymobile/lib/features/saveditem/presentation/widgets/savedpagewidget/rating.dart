import 'package:flutter/material.dart';

class RatingBar extends StatelessWidget {
  final double rating;
  final int starCount;
  final double size;
  final Color color;

  const RatingBar({
    super.key,
    required this.rating,
    this.starCount = 5,
    this.size = 24.0,
    this.color = const Color.fromARGB(255, 0, 0, 0),
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.start,
      children: [
        ...List.generate(starCount, (index) {
          return buildStar(context, index);
        }),
        Text(' $rating', style: TextStyle(fontSize: size * 0.8, color: color),)
      ]
    );
  }

  Widget buildStar(BuildContext context, int index) {
    IconData icon;
    if (rating >= index + 1) {
      icon = Icons.star;
    } else if (rating > index && rating < index + 1) {
      icon = Icons.star_half;
    } else {
      icon = Icons.star_border;
    }
    return Icon(icon, size: size, color: color);
  }
}
