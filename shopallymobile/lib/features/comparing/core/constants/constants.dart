// import 'package:shopally/features/comparing/data/model/product_model.dart';
// import '../../features/comparing/domain/Entity/product_entity.dart';

final String baseUrl = 'https://g6-shopally-3.onrender.com/api/v1';

// final response = {
//   "data": {
//     "products": [
//       {
//         "product": {
//           "id": "1005009254291888",
//           "title":
//               "QIALINO 15-Inch Laptop Sleeve: Waterproof, Shockproof, & Stylish",
//           "imageUrl":
//               "https://ae-pic-a1.aliexpress-media.com/kf/Sddc0076ca42d4a32962f8a617910314bn.jpg",
//           "aiMatchPercentage": 30,
//           "price": {
//             "etb": 3857.05,
//             "usd": 27.31,
//             "fxTimestamp": "2025-09-03T13:42:32.572898678Z",
//           },
//           "productRating": 0,
//           "deliveryEstimate": "",
//           "description":
//               "Protect your valuable 15-inch laptop with the QIALINO waterproof and shockproof sleeve. Crafted with high-quality materials, this carrying case offers superior protection against accidental bumps, scratches, and spills.",
//           "productSmallImageUrls": null,
//           "numberSold": 0,
//           "summaryBullets": [
//             "Waterproof Design: Keeps your laptop safe from spills and rain.",
//             "Shockproof Padding: Provides excellent protection against bumps and drops.",
//             "Convenient Carrying: Features a sturdy handle and detachable shoulder strap.",
//             "Sleek & Stylish: Complements your laptop with a professional look.",
//           ],
//           "deeplinkUrl":
//               "https://www.aliexpress.com/item/1005009254291888.html",
//           "taxRate": 0,
//           "discount": 20,
//         },
//         "synthesis": {
//           "pros": [
//             "Offers superior laptop protection with waterproof and shockproof design.",
//             "Provides convenient carrying options with a handle and shoulder strap.",
//             "Maintains a sleek and professional aesthetic.",
//           ],
//           "cons": [
//             "Higher price point compared to the sticker set.",
//             "No customer ratings available.",
//           ],
//           "isBestValue": false,
//           "features": {
//             "Delivery Speed": "Delivery speed is not specified.",
//             "Popularity & Demand": "No sales recorded.",
//             "Price & Value":
//                 "Significantly more expensive, but offers substantial protection.",
//             "Quality & Durability":
//                 "Designed with high-quality materials for long-lasting protection.",
//             "Seller Trust":
//                 "Seller trust is unknown as there's no seller score.",
//             "Unique Features":
//                 "Waterproof and shockproof capabilities with carrying options.",
//           },
//         },
//       },
//       {
//         "product": {
//           "id": "1005009710981318",
//           "title":
//               "54-Piece Anime Sticker Set - 'The Hundred Line' - Perfect for Laptops, Cars, and Gifts!",
//           "imageUrl":
//               "https://ae-pic-a1.aliexpress-media.com/kf/S0fac19d78b5c4c818f6c18c7329db5d4b.jpeg",
//           "aiMatchPercentage": 30,
//           "price": {
//             "etb": 927.9,
//             "usd": 6.57,
//             "fxTimestamp": "2025-09-03T13:42:32.636458648Z",
//           },
//           "productRating": 0,
//           "deliveryEstimate": "",
//           "description":
//               "Express your love for anime with this 54-piece sticker set featuring characters from 'The Hundred Line - Last Defense Academy'.",
//           "productSmallImageUrls": null,
//           "numberSold": 1,
//           "summaryBullets": [
//             "Unleash your inner anime fan with 54 unique stickers!",
//             "Transform ordinary items into personalized masterpieces!",
//             "Durable and vibrant, perfect for laptops, cars, and more!",
//             "Showcase your love for 'The Hundred Line - Last Defense Academy'!",
//           ],
//           "deeplinkUrl":
//               "https://www.aliexpress.com/item/1005009710981318.html",
//           "taxRate": 0,
//           "discount": 36,
//         },
//         "synthesis": {
//           "pros": [
//             "Affordable price point, significantly cheaper than the laptop sleeve.",
//             "Offers 54 unique anime stickers for personalization.",
//             "Suitable for various applications including laptops, cars, and gifts.",
//           ],
//           "cons": [
//             "Provides no functional protection like the laptop sleeve.",
//             "No customer ratings available.",
//           ],
//           "isBestValue": true,
//           "features": {
//             "Delivery Speed": "Delivery speed is not specified.",
//             "Popularity & Demand": "Low sales volume.",
//             "Price & Value":
//                 "Very affordable, offering a large number of stickers for a low price.",
//             "Quality & Durability":
//                 "Advertised as durable and vibrant, but lacks detailed material information.",
//             "Seller Trust":
//                 "Seller trust is unknown as there's no seller score.",
//             "Unique Features":
//                 "Large quantity of unique anime-themed stickers.",
//           },
//         },
//       },
//     ],
//     "overallComparison": {
//       "bestValueProduct":
//           "54-Piece Anime Sticker Set - 'The Hundred Line' - Perfect for Laptops, Cars, and Gifts!",
//       "bestValueLink": "https://www.aliexpress.com/item/1005009710981318.html",
//       "bestValuePrice": {
//         "etb": 927.9,
//         "usd": 6.57,
//         "fxTimestamp": "2025-09-03T13:42:32.636458648Z",
//       },
//       "keyHighlights": [
//         "The anime sticker set provides the best value due to its low price and large quantity of stickers.",
//         "The laptop sleeve offers superior protection but at a significantly higher cost.",
//       ],
//       "summary":
//           "The 54-Piece Anime Sticker Set offers the best value due to its significantly lower price and the quantity of stickers provided. The QIALINO Laptop Sleeve provides functional protection for a laptop, but at a much higher price point. The choice depends on the user's need for protection versus decorative personalization.",
//     },
//   },
//   "error": null,
// };


// final request = {
//   "products": [
//     {
//       "id": "1005009254291888",
//       "title": "QIALINO 15-Inch Laptop Sleeve: Waterproof, Shockproof, & Stylish",
//       "imageUrl": "https://ae-pic-a1.aliexpress-media.com/kf/Sddc0076ca42d4a32962f8a617910314bn.jpg",
//       "aiMatchPercentage": 30,
//       "price": {
//         "etb": 3857.05,
//         "usd": 27.31,
//         "fxTimestamp": "2025-09-03T13:42:32.572898678Z"
//       },
//       "productRating": 0,
//       "sellerScore": 0,
//       "deliveryEstimate": "",
//       "description": "Protect your valuable 15-inch laptop with the QIALINO waterproof and shockproof sleeve. Crafted with high-quality materials, this carrying case offers superior protection against accidental bumps, scratches, and spills.",
//       "customerHighlights": "Enjoy peace of mind knowing your laptop is shielded from bumps, spills, and scratches. The convenient handle and shoulder strap make it easy to carry.",
//       "customerReview": "I absolutely love this laptop sleeve! It's super lightweight but feels really durable. The waterproof feature is a lifesaver.",
//       "numberSold": 0,
//       "summaryBullets": [
//         "Waterproof Design: Keeps your laptop safe from spills and rain.",
//         "Shockproof Padding: Provides excellent protection against bumps and drops.",
//         "Convenient Carrying: Features a sturdy handle and detachable shoulder strap.",
//         "Sleek & Stylish: Complements your laptop with a professional look."
//       ],
//       "deeplinkUrl": "https://www.aliexpress.com/item/1005009254291888.html",
//       "taxRate": 0,
//       "discount": 20
//     },
//     {
//       "id": "1005009710981318",
//       "title": "54-Piece Anime Sticker Set - 'The Hundred Line' - Perfect for Laptops, Cars, and Gifts!",
//       "imageUrl": "https://ae-pic-a1.aliexpress-media.com/kf/S0fac19d78b5c4c818f6c18c7329db5d4b.jpeg",
//       "aiMatchPercentage": 30,
//       "price": {
//         "etb": 927.90,
//         "usd": 6.57,
//         "fxTimestamp": "2025-09-03T13:42:32.636458648Z"
//       },
//       "productRating": 0,
//       "sellerScore": 0,
//       "deliveryEstimate": "",
//       "description": "Express your love for anime with this 54-piece sticker set featuring characters from 'The Hundred Line - Last Defense Academy'.",
//       "customerHighlights": "Perfect for anime lovers, these stickers add a unique and personalized touch to your belongings.",
//       "customerReview": "These stickers are so cute! I used them to decorate my laptop and now it looks amazing.",
//       "numberSold": 1,
//       "summaryBullets": [
//         "Unleash your inner anime fan with 54 unique stickers!",
//         "Transform ordinary items into personalized masterpieces!",
//         "Durable and vibrant, perfect for laptops, cars, and more!",
//         "Showcase your love for 'The Hundred Line - Last Defense Academy'!"
//       ],
//       "deeplinkUrl": "https://www.aliexpress.com/item/1005009710981318.html",
//       "taxRate": 0,
//       "discount": 36
//     }
//   ]
// };

// final List<ProductEntity> products = (request['products'] as List)
//     .map((r) => ProductModel.fromJson(r).toEntity())
//     // .toList();