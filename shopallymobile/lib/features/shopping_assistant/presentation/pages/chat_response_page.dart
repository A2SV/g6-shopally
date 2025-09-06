import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../core/constants/ui_constants.dart';
import '../../domain/entities/product_entity.dart';
import '../bloc/chat_bloc.dart';
import '../widgets/product_link_box.dart';
import '../widgets/prompt_input.dart';

class ChatResponsePage extends StatefulWidget {
  final String prompt;
  const ChatResponsePage({super.key, required this.prompt});

  @override
  State<ChatResponsePage> createState() => _ChatResponsePageState();
}

class _ChatResponsePageState extends State<ChatResponsePage> {
  final List<List<ProductEntity>> products = [];
  final List<String> messages = [];
  final TextEditingController descriptionController = TextEditingController();

  @override
  void initState() {
    super.initState();
    messages.add(widget.prompt);
  }

  @override
  void dispose() {
    descriptionController.dispose();
    super.dispose();
  }

  Widget _myChatWidget(String text) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.end,
      children: [
        ConstrainedBox(
          constraints: BoxConstraints(
            maxWidth: MediaQuery.of(context).size.width * 0.7,
          ),
          child: Container(
            padding: const EdgeInsets.all(16),
            margin: const EdgeInsets.only(top: 8, right: 8, left: 16),
            decoration: const BoxDecoration(
              color: Color.fromARGB(255, 243, 236, 236),
              borderRadius: BorderRadius.only(
                topLeft: Radius.circular(20),
                bottomLeft: Radius.circular(20),
                bottomRight: Radius.circular(20),
              ),
            ),
            child: Text(text, style: const TextStyle(fontSize: 14)),
          ),
        ),
      ],
    );
  }

  Widget _responseBox(String text) {
    return Container(
      margin: const EdgeInsets.only(left: 16, right: 16),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        children: [
          Flexible(child: Text(text, style: const TextStyle(fontSize: 14))),
        ],
      ),
    );
  }

  Widget _loadingProducts() {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const SizedBox(height: 12),
        _responseBox('Preparing product suggestions...'),
        const SizedBox(height: 12),
        Container(
          margin: const EdgeInsets.symmetric(horizontal: 16),
          padding: const EdgeInsets.all(12),
          decoration: BoxDecoration(
            gradient: LinearGradient(
              colors: [Colors.white, Colors.grey.shade300],
            ),
            borderRadius: BorderRadius.circular(12),
          ),
          child: const Row(
            children: [
              Icon(Icons.link, color: Colors.blue),
              SizedBox(width: 8),
              Expanded(child: Text('Searching products...')),
              SizedBox(
                width: 16,
                height: 16,
                child: CircularProgressIndicator(strokeWidth: 2),
              ),
            ],
          ),
        ),
        const SizedBox(height: 12),
      ],
    );
  }

  Widget _productsSection(int index) {
    final query = messages[index];
    final list = products[index];
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const SizedBox(height: 12),
        _responseBox('I am searching for $query related content for you...'),
        const SizedBox(height: 12),
        ProductLinkBox(products: list, text: query),
        const SizedBox(height: 12),
        _responseBox(
          'Found ${list.length} $query related option(s). Feel free to explore them!',
        ),
        const SizedBox(height: 12),
      ],
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.backgroundTop,
      appBar: AppBar(
        backgroundColor: AppColors.backgroundTop,
        leading: BackButton(
          onPressed: () => Navigator.pushNamed(context, '/chat'),
        ),
        title: const  Text('Chat'),
      ),
      body: Column(
        children: [
          const SizedBox(height: 20),
          Expanded(
            child: BlocConsumer<ChatBloc, ChatState>(
              listener: (context, state) {
                if (state is ProductsLoadedState) {
                  setState(() {
                    products.add(state.products);
                  });
                }
              },
              builder: (context, state) {
                if (state is ResponseErrorState) {
                  return Padding(
                    padding: const EdgeInsets.all(16),
                    child: Text(
                      'Error: ${state.message}',
                      style: const TextStyle(color: Colors.red),
                    ),
                  );
                }
                return SingleChildScrollView(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      ListView.builder(
                        shrinkWrap: true,
                        physics: const NeverScrollableScrollPhysics(),
                        itemCount: messages.length,
                        itemBuilder: (context, index) {
                          final hasProducts = index < products.length;
                          return Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              const SizedBox(height: 12),
                              _myChatWidget(messages[index]),
                              hasProducts
                                  ? _productsSection(index)
                                  : _loadingProducts(),
                            ],
                          );
                        },
                      ),
                    ],
                  ),
                );
              },
            ),
          ),
          PromptInput(
            onSubmit: () {
              final text = descriptionController.text.trim();
              if (text.isEmpty) return;
              setState(() {
                messages.add(text);
              });
              context.read<ChatBloc>().add(SendPromptRequested(text));
              descriptionController.clear();
            },
            descriptionController: descriptionController,
          ),
        ],
      ),
    );
  }
}
