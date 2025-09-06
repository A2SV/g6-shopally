import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shopallymobile/core/localization/localization_store.dart';
import '../bloc/chat_bloc.dart';

class PromptInput extends StatefulWidget {
  final VoidCallback onSubmit;
  final TextEditingController descriptionController;

  const PromptInput({
    super.key,
    required this.onSubmit,
    required this.descriptionController,
  });

  @override
  State<PromptInput> createState() => _PromptInputState();
}

class _PromptInputState extends State<PromptInput> {
  @override
  void initState() {
    super.initState();
  }

  @override
  void dispose() {
    widget.descriptionController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      constraints: BoxConstraints(minHeight: 60),
      padding: const EdgeInsets.only(
        left: 5.0,
        right: 5.0,
        top: 5.0,
        bottom: 5.0,
      ),
      margin: const EdgeInsets.all(15.0),
      decoration: BoxDecoration(
        color: Theme.of(context).cardColor,
        borderRadius: BorderRadius.circular(30.0),
        boxShadow: [
          BoxShadow(
            color: Colors.black.withOpacity(
              Theme.of(context).brightness == Brightness.dark ? 0.2 : 0.06,
            ),
            blurRadius: 6,
            offset: const Offset(0, 2),
          ),
        ],
      ),
      child: Row(
        children: [
          Expanded(
            child: TextField(
              keyboardType: TextInputType.multiline,
              minLines: 1,
              controller: widget.descriptionController,
              onChanged: (value) {
                // Trigger rebuild to update send button state
                setState(() {});
              },
              maxLines: null,
              style: const TextStyle(fontSize: 14.0),
              decoration: InputDecoration(
                hintText: getText('describe_your_shopping_needs'),
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(8.0),
                  borderSide: BorderSide.none,
                ),
              ),
            ),
          ),
          const SizedBox(width: 10),
          BlocBuilder<ChatBloc, ChatState>(
            builder: (context, state) {
              final primaryTextColor =
                  Theme.of(context).textTheme.bodyLarge?.color ?? Colors.black;
              final active =
                  widget.descriptionController.text.isNotEmpty ||
                  state is LoadingState ||
                  state is ResponseLoadingState;
              return ElevatedButton(
                style: ElevatedButton.styleFrom(
                  shape: const CircleBorder(),
                  padding: const EdgeInsets.all(0),
                  backgroundColor: active
                      ? primaryTextColor
                      : Theme.of(context).cardColor,
                ),
                onPressed: active
                    ? () {
                        widget.onSubmit();
                      }
                    : null,
                child: BlocBuilder<ChatBloc, ChatState>(
                  builder: (context, state) {
                    final iconColor = active
                        ? (Theme.of(context).brightness == Brightness.dark
                              ? Colors.black
                              : Colors.white)
                        : primaryTextColor;
                    if (state is LoadingState ||
                        state is ResponseLoadingState) {
                      return Icon(Icons.square, size: 14, color: iconColor);
                    }
                    return Icon(Icons.send, size: 14, color: iconColor);
                  },
                ),
              );
            },
          ),
        ],
      ),
    );
  }
}
