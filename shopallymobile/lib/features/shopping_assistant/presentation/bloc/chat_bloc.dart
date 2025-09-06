import 'package:bloc/bloc.dart';
import 'package:meta/meta.dart';

import '../../domain/entities/product_entity.dart';
import '../../domain/usecases/send_prompt.dart';

part 'chat_event.dart';
part 'chat_state.dart';

class ChatBloc extends Bloc<ChatEvent, ChatState> {
  final SendPrompt sendPrompt;

  ChatBloc({required this.sendPrompt}) : super(ChatInitial()) {
    on<ChatEvent>((event, emit) {
      
    });
    on<SendPromptRequested>(_onSendPromptRequested);
  }
  Future<void> _onSendPromptRequested(SendPromptRequested event, Emitter<ChatState> emit) async {
    emit(ResponseLoadingState());
    final result = await sendPrompt(event.prompt);
    result.fold(
      (failure) {
        // Handle failure state if needed
        emit(ResponseErrorState(failure.message)); // or some error state
      },
      (products) {
        emit(ProductsLoadedState(products: products));
      },
    );
  }
}
