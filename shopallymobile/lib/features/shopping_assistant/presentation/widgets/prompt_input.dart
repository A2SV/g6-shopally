import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../core/constants/ui_constants.dart';
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
        color: Colors.white,
        borderRadius: BorderRadius.circular(30.0),
        boxShadow: [
          for (int i = 0; i < AppColors.inputGlowGradient.length; i++)
            BoxShadow(
              color: AppColors.inputGlowGradient[i].withValues(
                alpha: 0.20 + i * 0.05,
              ),
              blurRadius: 1 + i * 1.0,
              spreadRadius: 0,
              offset: Offset(0, 0.004 + i * .5),
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
                hintText: 'Describe your shopping needs...',
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
              return ElevatedButton(
                style: ElevatedButton.styleFrom(
                  shape: const CircleBorder(),
                  padding: const EdgeInsets.all(0),
                  backgroundColor: widget.descriptionController.text.isNotEmpty
                      ? Colors.black
                      : Colors.white,
                ),
                onPressed: (widget.descriptionController.text.isNotEmpty || state is LoadingState)
                    ? () {
                        widget.onSubmit();
                      }
                    : null,
                child: BlocBuilder<ChatBloc, ChatState>(
                  builder: (context, state) {
                    if (state is LoadingState || state is ResponseLoadingState) {
                      return Icon(
                        Icons.square,
                        size: 14,
                        color: widget.descriptionController.text.isNotEmpty
                            ? Colors.white
                            : Colors.black,
                      );
                    }
                    return Icon(
                      Icons.send,
                      size: 14,
                      color: widget.descriptionController.text.isNotEmpty
                          ? Colors.white
                          : Colors.black,
                    );
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
