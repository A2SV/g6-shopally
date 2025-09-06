import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shopallymobile/features/saveditem/presentation/bloc/bloc/saved_product_bloc.dart';
import 'package:shopallymobile/features/saveditem/presentation/widgets/savedpagewidget/singel_product_info.dart';

class Savedpage extends StatefulWidget {
  const Savedpage({super.key});

  @override
  State<Savedpage> createState() => _SavedpageState();
}

class _SavedpageState extends State<Savedpage> {
  // Cache last successfully loaded products to avoid UI changes on refresh.
  List<dynamic> _productsCache = [];
  bool _isLoading = false;

  @override
  void initState() {
    super.initState();
    _isLoading = true; // show overlay until first data arrives
    context.read<SavedProductBloc>().add(LoadSavedProductsEvent());
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: BlocConsumer<SavedProductBloc, SavedProductState>(
        listener: (context, state) {
          if (state is SavedProductLoading) {
            setState(() => _isLoading = true);
          } else if (state is SavedProductLoaded) {
            setState(() {
              _productsCache = state.products;
              _isLoading = false;
            });
          } else if (state is SavedProductError ||
              state is RemoveProductError ||
              state is SaveProductError ||
              state is AddToCompareError) {
            setState(() => _isLoading = false);
            final msg = state is SavedProductError
                ? state.message
                : state is RemoveProductError
                    ? state.message
                    : state is SaveProductError
                        ? state.message
                        : state is AddToCompareError
                            ? state.message
                            : 'Something went wrong';
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(content: Text(msg)),
            );
          }
        },
        builder: (context, state) {
          final savedProducts = _productsCache;

          return Stack(
            children: [
              CustomScrollView(
                slivers: [
                  SliverAppBar(
                    expandedHeight: 150,
                    pinned: true,
                    floating: true,
                    shape: const RoundedRectangleBorder(
                      borderRadius: BorderRadius.vertical(
                        bottom: Radius.circular(30),
                      ),
                    ),
                    flexibleSpace: FlexibleSpaceBar(
                      centerTitle: true,
                      title: Text(
                        savedProducts.isNotEmpty
                            ? '${savedProducts.length} Saved Item'
                            : '',
                        style: const TextStyle(fontSize: 16.0),
                      ),
                      background: Container(
                        decoration: BoxDecoration(
                          borderRadius: const BorderRadius.vertical(
                            bottom: Radius.circular(30),
                          ),
                          gradient: LinearGradient(
                            begin: Alignment.topLeft,
                            end: Alignment.bottomRight,
                            colors: [
                              Theme.of(context).primaryColor,
                              Theme.of(context).colorScheme.secondary,
                            ],
                          ),
                        ),
                      ),
                    ),
                  ),
                  if (savedProducts.isEmpty)
                    const SliverFillRemaining(
                      child: Center(
                        child: Text(
                          'No saved items yet.',
                          style: TextStyle(fontSize: 18, color: Colors.grey),
                        ),
                      ),
                    )
                  else
                    SliverPadding(
                      padding: const EdgeInsets.all(8.0),
                      sliver: SliverGrid(
                        gridDelegate:
                            const SliverGridDelegateWithFixedCrossAxisCount(
                          crossAxisCount: 2,
                          crossAxisSpacing: 2.0,
                          mainAxisSpacing: 8.0,
                          childAspectRatio: 0.55,
                        ),
                        delegate: SliverChildBuilderDelegate(
                          (context, index) {
                            final product = savedProducts[index];
                            return GestureDetector(
                              key: ValueKey(product.id), // stable identity
                              onTap: () {
                                // Navigator.push(
                                //   context,
                                //   MaterialPageRoute(
                                //     builder: (context) => DetailPage(product: product),
                                //   ),
                                // );
                              },
                              child: SingleProductInfo(
                                id: product.id,
                                title: product.title,
                                price: product.price,
                                minOrder: product.minOrder,
                                rating: product.rating,
                                image: product.imageUrl,
                              ),
                            );
                          },
                          childCount: savedProducts.length,
                        ),
                      ),
                    ),
                ],
              ),
              if (_isLoading)
                const Positioned.fill(
                  child: IgnorePointer(
                    child: Center(child: CircularProgressIndicator()),
                  ),
                ),
            ],
          );
        },
      ),
    );
  }
}