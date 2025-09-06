part of 'compare_bloc.dart';

@immutable
sealed class CompareState {}

final class CompareInitial extends CompareState {}

final class CompareLoading extends CompareState {}

final class ComparingProductLoaded extends CompareState {
  final List<ProductEntity> products;
  ComparingProductLoaded(this.products);
}

final class CompareError extends CompareState {
  final String message;
  CompareError(this.message);
}

final class ComparisonEmpty extends CompareState {}

final class ComparisonResult extends CompareState {
  final ComparisonResultEntity comparisons;
  ComparisonResult(this.comparisons);
}
