import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'package:shopallymobile/core/localization/localization_store.dart';
import 'package:shopallymobile/features/comparing/presentation/bloc/compare_bloc.dart';
import 'package:shopallymobile/features/comparing/presentation/pages/comparison_result_page.dart';
import 'package:shopallymobile/features/comparing/presentation/widget/product_card.dart';

import '../widget/error_page.dart';
import '../widget/recommendation.dart';

class ProductsForComparing extends StatefulWidget {
  const ProductsForComparing({super.key});

  @override
  State<ProductsForComparing> createState() => _ProductsForComparingState();
}

class _ProductsForComparingState extends State<ProductsForComparing> {
  @override
  Widget build(BuildContext context) {
    return BlocConsumer<CompareBloc, CompareState>(
      listener: (context, state) {
        if (state is CompareError) {
          WidgetsBinding.instance.addPostFrameCallback((_) {
            showDialog(
              context: context,
              barrierDismissible: false,
              builder: (_) => ErrorPage(
                message: state.message,
                onDismiss: () {
                  context.read<CompareBloc>().add(
                    FetchProductForComparisonEvent(),
                  );
                },
              ),
            );
          });
        } else if (state is ComparisonResult) {
          Navigator.push(
            context,
            MaterialPageRoute(
              builder: (context) => ComparisonResultPage(
                comparisonResultEntity: state.comparisons,
              ),
            ),
          );
        } else if (state is CompareInitial) {
          context.read<CompareBloc>().add(FetchProductForComparisonEvent());
        }
      },
      builder: (context, state) {
        return BlocBuilder<CompareBloc, CompareState>(
          builder: (context, state) {
            return _buildContent(state);
          },
        );
      },
    );
  }

  Widget loadingState() {
    return Container(
      width: double.infinity,
      height: double.infinity,
      child: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            // Animated shopping bag
            RotationTransition(
              turns: AlwaysStoppedAnimation(0.1),
              child: Icon(
                Icons.shopping_bag,
                size: 80,
                color: Color.fromRGBO(255, 211, 0, 1),
              ),
            ),
            SizedBox(height: 20),

            // Loading text with fade animation
            // FadeTransition(
            //   opacity: AlwaysStoppedAnimation(0.7),
            //   child: Text(
            //     getText('Loading ...'),
            //     style: TextStyle(
            //       fontSize: 18,
            //       fontWeight: FontWeight.w500,
            //       color: Colors.grey[600],
            //     ),
            //   ),
            // ),
            // SizedBox(height: 20),

            // Progress bar
            SizedBox(
              width: 200,
              child: LinearProgressIndicator(
                backgroundColor: Colors.grey[200],
                valueColor: AlwaysStoppedAnimation<Color>(
                  Color.fromRGBO(255, 211, 0, 1),
                ),
                borderRadius: BorderRadius.circular(10),
                minHeight: 6,
              ),
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildContent(CompareState state) {
    if (state is ComparingProductLoaded) {
      return ProductCard(productEntities: state.products);
    } else if (state is CompareLoading) {
      return loadingState();
    } else if (state is ComparisonResult) {
      return ComparisonResultPage(comparisonResultEntity: state.comparisons);
    } else if (state is CompareError) {
      WidgetsBinding.instance.addPostFrameCallback((_) {
        showDialog(
          context: context,
          barrierDismissible: false,
          builder: (_) => ErrorPage(
            message: state.message,
            onDismiss: () {
              context.read<CompareBloc>().add(FetchProductForComparisonEvent());
            },
          ),
        );
      });
      return Container(); // Empty container behind
    } else if (state is ComparisonEmpty) {
      return Center(
        child: Padding(
          padding: const EdgeInsets.all(24.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Icon(
                Icons.compare_arrows_rounded,
                size: 100,
                color: Colors.grey.shade400,
              ),
              const SizedBox(height: 20),
              Text(
                getText("No Products to Compare"),
                style: TextStyle(
                  fontSize: 22,
                  fontWeight: FontWeight.bold,
                  color: Color.fromRGBO(38, 43, 50, 1),
                ),
              ),
              const SizedBox(height: 8),
              const Text(
                "Add products to your comparison list to get started.",
                textAlign: TextAlign.center,
                style: TextStyle(
                  fontSize: 16,
                  color: Color.fromRGBO(117, 123, 129, 1),
                ),
              ),
              const SizedBox(height: 24),
              ElevatedButton(
                onPressed: () {
                  // Navigate to products page
                },
                style: ElevatedButton.styleFrom(
                  backgroundColor: Colors.blue,
                  foregroundColor: Colors.white,
                  padding: const EdgeInsets.symmetric(
                    horizontal: 24,
                    vertical: 12,
                  ),
                ),
                child: const Text('Browse Products'),
              ),
            ],
          ),
        ),
      );
    } else {
      return Container();
    }
  }
}
