import 'package:flutter/material.dart';
import 'package:shopallymobile/core/localization/localization_store.dart';

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
         Padding(
          padding: EdgeInsets.symmetric(horizontal: 16.0),
          child: Text(
            getText('suggestion'),
            style: TextStyle(
              fontSize: 18, // slightly smaller title if desired
              fontWeight: FontWeight.bold,
              color: Theme.of(context).textTheme.bodyLarge?.color,
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
              _buildSuggestionCard(getText('find_me_a_red_dress')),
              _buildSuggestionCard(getText('best_budget_laptops')),
              _buildSuggestionCard(getText('top_rated_headphones')),
              _buildSuggestionCard(getText('affordable_smartphones')),
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
          color: Theme.of(context).cardColor,
          borderRadius: BorderRadius.circular(18.0),
          boxShadow: [
            BoxShadow(
              color: Colors.black.withOpacity(
                Theme.of(context).brightness == Brightness.dark ? 0.2 : 0.08,
              ),
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
                  style: TextStyle(
                    color: Theme.of(context).textTheme.bodyLarge?.color,
                    fontSize: 16,
                    fontWeight: FontWeight.w500,
                  ),
                ),
              ),
              const SizedBox(height: 2),
              Text(
                getText('suggested'),
                style: TextStyle(
                  color: Theme.of(context).textTheme.bodyMedium?.color?.withOpacity(0.75),
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
