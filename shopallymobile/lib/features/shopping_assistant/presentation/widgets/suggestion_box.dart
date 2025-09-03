import 'package:flutter/material.dart';

import '../../../../core/constants/ui_constants.dart';

class SuggestionBox extends StatefulWidget {
  final TextEditingController descriptionController;
  const SuggestionBox({super.key, required this.descriptionController});

  @override
  State<SuggestionBox> createState() => _SuggestionBoxState();
}

class _SuggestionBoxState extends State<SuggestionBox> {
  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const Padding(
          padding: EdgeInsets.symmetric(horizontal: 16.0),
          child: Text(
            'Suggestions',
            style: TextStyle(
              fontSize: 18, // slightly smaller title if desired
              fontWeight: FontWeight.bold,
            ),
          ),
        ),
        const SizedBox(height: 8),
        SizedBox(
          height: 64, // reduced height (was 100)
          child: ListView(
            scrollDirection: Axis.horizontal,
            padding: const EdgeInsets.symmetric(horizontal: 16.0),
            children: [
              _buildSuggestionCard('Find me a red dress'),
              _buildSuggestionCard('Best budget laptops'),
              _buildSuggestionCard('Top-rated headphones'),
              _buildSuggestionCard('Affordable smartphones'),
            ],
          ),
        ),
      ],
    );
  }

  Widget _buildSuggestionCard(String text) {
    return GestureDetector(
      onTap: () {
        setState(() {
          widget.descriptionController.text = text;
        });
      },
      child: Container(
        margin: const EdgeInsets.only(right: 10.0, bottom: 6.0),
        padding: const EdgeInsets.symmetric(horizontal: 12.0, vertical: 6.0),
        constraints: const BoxConstraints(minWidth: 110),
        decoration: BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.circular(18.0),
          boxShadow: [
            BoxShadow(
              color: Colors.grey.withOpacity(0.18),
              blurRadius: 4,
              offset: const Offset(0, 2),
            ),
          ],
        ),
        child: SizedBox(
          height: 48, // keep internal height consistent
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Flexible(
                child: Text(
                  text,
                  softWrap: true,
                  maxLines: 2,
                  overflow: TextOverflow.ellipsis,
                  style: const TextStyle(
                    color: AppColors.textPrimary,
                    fontSize: 16,
                    fontWeight: FontWeight.w500,
                  ),
                ),
              ),
              const SizedBox(height: 2),
              Text(
                'Suggested',
                style: TextStyle(
                  color: AppColors.textSecondary,
                  fontSize: 12,
                  fontWeight: FontWeight.w400,
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
