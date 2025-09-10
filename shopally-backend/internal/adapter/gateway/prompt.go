package gateway

import (
	"fmt"

	"github.com/shopally-ai/pkg/domain"
)

func intentParsingPrompt(userPrompt string) string {
	return fmt.Sprintf(`STRICT INSTRUCTIONS: OUTPUT ONLY RAW JSON, NO OTHER TEXT, NO EXPLANATIONS, NO CODE BLOCKS.

			You are an advanced multi-language e-commerce intent parser. Your task is to normalize user queries into structured JSON for product search.

			## CRITICAL RULES:
			1. OUTPUT ONLY PURE JSON.
			2. DETECT CURRENCY MENTIONED IN QUERY:
			- If query mentions "USD", "$", or "dollars" → set is_etb=false.
			- If query mentions "ETB", "birr", "ብር" or no currency → set is_etb=true.
			3. EXTRACT PRODUCT-ONLY KEYWORDS:
			- Keep ONLY product identifiers: main product type, brand names, model numbers, usage/function, and gender/audience.
			- REMOVE all filler words (AliExpress-style marketing noise): "new", "2024", "hot", "latest", "luxury", "fashion", "high quality", "original", "authentic", "offer", "trending", "classic", "design", "style", "durable", "comfortable", "beautiful", "discount", "sale", "free shipping".
			- DO NOT include generic adjectives or promotional words — only the core product intent.
			- STEM WORDS TO BASE FORM (e.g., "running" → "run", "shoes" → "shoe").
			- TRANSLATE PRODUCT TERMS TO ENGLISH.
			- MERGE EVERYTHING into a single space-separated string → keywords.
			- INCLUDE gender (male, female, unisex, kid) IF SPECIFIED.
			- PRESERVE brand names and model numbers (e.g., "Nike AirMax 270" → "nike airmax 270").
			4. CONVERT NUMBER WORDS (English & Amharic) TO DIGITS.
			5. PRESERVE ORIGINAL BUDGET PHRASE as given by the user.
			6. NORMALIZE QUERY INTO GENERIC CLASS → query_class.
			- The generic class represents only the core product intent.
			- Example: "cheap gaming laptop under 1000 USD" → "gaming laptop".
			- Example: "red nike running shoes" → "nike run shoe".
			7. DETECT LANGUAGE USED IN PROMPT:
			- If written in Amharic script or Latin transliteration → language = "am".
			- Else → language = "en".
			8. ASSIGN CATEGORY_ID:
			- Choose the closest matching category_id from the full labeled mapping list (provided separately in system memory).
			- The choice must be fully dependent on the user prompt.
			- Always return exactly ONE best-fitting category ID as a string in "category_ids".

			## FINAL JSON STRUCTURE:
			{
			"keywords": "string",
			"min_sale_price": number|null,
			"max_sale_price": number|null,
			"original_budget": "string|null",
			"delivery_days": number|null,
			"ship_to_country": "ET",
			"target_currency": "USD",
			"target_language": "en",
			"is_etb": boolean,
			"query_class": "string",
			"language": "am" | "en",
			"category_ids": "string"
			}
			## EXAMPLES:

			"user query: red nike running shoes under 3000 birr" ->
			{
			"keywords": "red nike run shoe sneaker sport footwear male",
			"min_sale_price": null,
			"max_sale_price": 3000,
			"original_budget": "under 3000 birr",
			"delivery_days": null,
			"ship_to_country": "ET",
			"target_currency": "USD",
			"target_language": "en",
			"is_etb": true,
			"query_class": "nike run shoe"
			"language": "en"
			}

			"user query: cheap gaming laptop around one thousand dollars" ->
			{
			"keywords": "cheap game laptop notebook computer pc electronic",
			"min_sale_price": null,
			"max_sale_price": 1000,
			"original_budget": "around one thousand dollars",
			"delivery_days": null,
			"ship_to_country": "ET",
			"target_currency": "USD",
			"target_language": "en",
			"is_etb": false,
			"query_class": "game laptop"
			"language": "en"
			}

			"user query: ነጭ ቀሚስ የወንድ ከአምስት መቶ ብር በታች" ->
			{
			"keywords": "white shirt men male clothing apparel fashion",
			"min_sale_price": null,
			"max_sale_price": 500,
			"original_budget": "ከአምስት መቶ ብር በታች",
			"delivery_days": null,
			"ship_to_country": "ET",
			"target_currency": "USD",
			"target_language": "en",
			"is_etb": true,
			"query_class": "shirt men"
			"language: "am"
			}

			"user query: samsung galaxy s23 ultra phone under 700 dollars" ->
			{
			"keywords": "samsung galaxy s23 ultra phone smartphone mobile",
			"min_sale_price": null,
			"max_sale_price": 700,
			"original_budget": "under 700 dollars",
			"delivery_days": null,
			"ship_to_country": "ET",
			"target_currency": "USD",
			"target_language": "en",
			"is_etb": false,
			"query_class": "samsung galaxy s23 ultra phone"
			"language": "en"
			}

			"user query: 2024 New Luxury Fashion Men Casual Sport Running Shoes High Quality" ->
			{
			"keywords": "men run shoe sneaker sport",
			"min_sale_price": null,
			"max_sale_price": null,
			"original_budget": null,
			"delivery_days": null,
			"ship_to_country": "ET",
			"target_currency": "USD",
			"target_language": "en",
			"is_etb": true,
			"query_class": "run shoe"
			"language" "en"
}


			## CATEGORY MAPPING LIST:
			Home Appliances: 6
			Computer and Office: 7
			Home Improvement: 13
			Home & Garden: 15
			Sports & Entertainment: 18
			Education & Office Supplies: 21
			Toys & Hobbies: 26
			Security & Protection: 30
			Automobiles & Motorcycles: 34
			Lights & Lighting: 39
			Consumer Electronics: 44
			Beauty & Health: 66
			Shoes: 322
			Luggage & Bags: 1524
			Electronic Components & Supplies: 502
			Tools: 1420
			Mother & Kids: 1501
			Furniture: 1503
			Jewelry & Accessories: 1509
			Watches: 1511
			Hair Extensions and Wigs: 200002489
			Virtual Goods: 205965401
			Novelty and Special Use: 200000875
			Weddings and Events: 100003235
			Women’s Clothing: 100003109
			Men’s Clothing: 100003070
			Apparel Accessories: 205776616
			Underwear and Sleepwears: 205779615
			Cellphones & Telecommunications: 509
			Household Appliances: 200214052
			Personal Care Appliances: 200214073
			Commercial Appliances: 200217027
			Major Appliances: 200217594
			Home Appliance Parts: 100000016
			Kitchen Appliances: 100000011
			Storage Devices: 200215304
			Laptops: 702
			Servers: 703
			Demo Board and Accessories: 200216762
			Desktops: 200216675
			Tablets: 200216621
			Computer Cables and Connectors: 200216562
			Office Software: 205342002
			Mini PC: 70803003
			Computer Peripherals: 200002342
			Tablet Accessories: 200002361
			Networking: 200002320
			Computer Components: 200002319
			Device Cleaners: 708022
			Office Electronics: 200004720
			Industrial Computer and Accessories: 100005089
			Mouse and Keyboards: 100005085
			Laptop Accessories: 100005063
			Laptop Parts: 205848303
			Tablet Parts: 205845408
			Electrical Equipments and Supplies: 5
			Hardware: 42
			Kitchen Fixtures: 200215252
			Plumbing: 200217293
			Painting Supplies and Wall Treatments: 200217241
			Family Intelligence System: 200217718
			Building Supplies: 200003230
			Bathroom Fixtures: 100006479
			Garden Supplies: 125
			Household Merchandises: 200215281
			Home Textile: 405
			Home Storage and Organization: 1541
			Home Decor: 3710
			Arts, Crafts and Sewing: 200154001
			Kitchen, Dining and Bar: 200002086
			Household Cleaning: 200003136
			Bathroom Products: 100004814
			Pet Products: 100006206
			Sports Accessories: 200214370
			Sports Bags: 200217620
			Team Sports: 200094001
			Sports Clothing: 200001095
			Swimming: 200001115
			Cycling: 200003570
			Sneakers: 200005276
			Running: 200005156
			Roller Skates, Skateboards and Scooters: 200005143
			Bowling: 200005102
			Entertainment: 200005101
			Racquet Sports: 200005059
			Golf: 100005322
			Fitness and Body Building: 100005259
			Other Sports and Entertainment: 100005663
			Skiing and Snowboarding: 100005599
			Water Sports: 100005575
			Shooting: 100005479
			Horse Racing: 100005460
			Hunting: 100005471
			Fishing: 100005444
			Camping and Hiking: 100005433
			Musical Instruments: 100005383
			Paper: 2112
			Desk Accessories and Organizer: 211106
			Art Supplies: 211111
			Presentation Supplies: 212002
			Books and Magazines: 205954820
			Stationery Sticker: 205953922
			Mail and Shipping Supplies: 200003238
			Writing and Correction Supplies: 200003196
			Calendars, Planners and Cards: 200003198
			Labels, Indexes and Stamps: 200003197
			Filing Products: 100003804
			Notebooks and Writing Pads: 100003745
			School and Educational Supplies: 100005094
			Tapes, Adhesives and Fasteners: 100003836
			Office Binding Supplies: 100003809
			Cutting Supplies: 100003819
			Stress Relief Toy: 200216936
			Pools and Water Fun: 200218444
			Popular Toys: 200218404
			Hobby and Collectibles: 200218343
			Stuffed Animals and Plush: 200218367
			Arts and Crafts, DIY toys: 200218357
			High Tech Toys: 200218333
			Kid’s Party: 200218291
			Building and Construction Toys: 200218269
			Play Vehicles and Models: 206081903
			Ride On Toys: 206089103
			Action and Toy Figures: 206086301
			Puzzles and Games: 200003226
			Dolls and Accessories: 200003225
			Novelty and Gag Toys: 200002636
			Model Building: 200002633
			Remote Control Toys: 200002639
			Diecasts and Toy Vehicles: 100001663
			Baby and Toddler Toys: 100001622
			Pretend Play: 100001624
			Outdoor Fun and Sports: 100001623
			Electronic Toys: 100001629
			Classic Toys: 100001626
			Learning and Education: 100001625
			Security Alarm: 200215432
			Building Automation: 200215427
			Smart Card System: 200215424
			Door Intercom: 200215419
			Workplace Safety Supplies: 3007
			Fire Protection: 3009
			Video Surveillance: 3011
			Safes: 3012
			Self Defense Supplies: 3019
			Access Control: 3030
			Public Broadcasting: 205718021
			Roadway Safety: 200216744
			Transmission and Cables: 200216754
			IoT Devices: 205662015
			Security Inspection Device: 205676017
			Lightning Protection: 300912
			Emergency Kits: 200003251
			Car Lights: 200216084
			Car Repair Tools: 200216017
			ATV, RV, Boat and Other Vehicle: 200214451
			Travel and Roadway Product: 200217080
			Car Wash and Maintenance: 200217078
			Motorcycle Accessories and Parts: 200000408
			Car Electronics: 200000369
			Auto Replacement Parts: 200000191
			Exterior Accessories: 200004620
			Interior Accessories: 200004619
			Lighting Accessories: 530
			Holiday Lighting: 200216091
			Ceiling Lights and Fans: 1504
			Lamps and Shades: 200214033
			Special Engineering Lighting: 200217736
			Under Cabinet Lights: 200217706
			Outdoor Lighting: 150401
			Light Bulbs: 150402
			Novelty Lighting: 200002283
			Professional Lighting: 200003210
			Commercial Lighting: 200003009
			LED Lamps: 200003575
			Book Lights: 39050501
			Night Lights: 39050508
			Vanity Lights: 201976010
			LED Lighting: 390501
			Portable Lighting: 390503
			VR/AR Devices: 200215272
			Speakers: 200217534
			Sports and Action Video Cameras: 200216648
			Earphones and Headphones: 200216623
			360° Video Cameras and Accessories: 200216592
			Home Electronic Accessories: 200216598
			Power Source: 200218547
			Live Equipment: 200218521
			HIFI Devices: 200217800
			Robot: 200217794
			Wearable Devices: 200084019
			Video Games: 200002396
			Camera and Photo: 200002395
			Accessories and Parts: 200002394
			Portable Audio and Video: 200002398
			Home Audio and Video: 200002397
			Electronic Cigarettes: 200005280
			Smart Electronics: 200010196
			Sanitary Paper: 1513
			Oral Hygiene: 3305
			Skin Care: 3306
			Makeup: 660103
			Shaving and Hair Removal: 660302
			Beauty Equipment: 205820203
			Health Care: 200002496
			Bath and Shower: 200002444
			Hair Care and Styling: 200002458
			Fragrances and Deodorants: 200002454
			Tattoo and Body Art: 200003551
			Sex Products: 200003045
			Tools and Accessories: 200002569
			Nails Art and Tools: 200002547
			Men’s Grooming: 201902009
			Skin Care Tools: 205778203
			Beauty Essentials: 200074001
			Men’s Vulcanize Shoes: 200002253

			Women’s Boots: 200216407
			Women’s Flats: 200002155
			Women’s Pumps: 200002161
			Men’s Casual Shoes: 200002136
			Women’s Vulcanize Shoes: 200002164
			Men’s Boots: 200216391
			Shoe Accessories: 200002124
			Women’s Shoes: 100001606
			Men’s Shoes: 100001615
			Coin Purses and Holders: 3803
			Luggage and Travel Bags: 152404
			Bag Parts and Accessories: 152409
			Backpacks: 152401
			Wallets: 152405
			Functional Bags: 200068019
			Kids and Baby’s Bags: 200066014
			Men’s Bags: 200010057
			Women’s Bags: 200010063
			Electronic Data Systems: 504
			Electronics Stocks: 515
			Active Components: 4001
			Electronics Production Machinery: 4002
			Electronic Accessories and Supplies: 4003
			Optoelectronic Displays: 4004
			Passive Components: 4005
			Other Electronic Components: 4099
			EL Products: 150412
			Power Tools: 1417
			Welding and Soldering Supplies: 1427
			Abrasives: 1428
			Woodworking Machinery and Parts: 1431
			Welding Equipment: 1440
			Measurement and Analysis Instruments: 1537
			Abrasive Tools: 4204
			Riveter Guns: 200216862
			Hand and Power Tool Accessories: 200218051
			Lifting Tools and Accessories: 200218021
			Hand Tools: 142003
			Tool Parts: 142001
			Construction Tools: 142016
			Garden Tools: 12503
			Tool Sets: 100006919
			Tool Organizers: 100006925
			Machine Tools and Accessories: 100006799
			Baby Stroller and Accessories: 200217552
			Toilet Training: 200217567
			Car Seats and Accessories: 200217523
			Baby Food: 200218586
			Pregnancy and Maternity: 200217581
			Baby Souvenirs: 200217580
			Baby Furniture: 200217573
			Matching Family Outfits: 200166001
			Kids and Baby Accessories: 205870202
			Children’s Shoes: 32212
			Nappy Changing: 200002433
			Baby Shoes: 200002101
			Girls’ Baby Clothing: 200000567
			Boys’ Baby Clothing: 200000528
			Feeding: 200003595
			Activity and Gear: 200003594
			Safety Equipment: 200003592
			Baby Care: 100001118
			Boys’ Clothing: 100003186
			Girls’ Clothing: 100003199
			Baby Bedding: 100002964
			Cafe Furniture: 200216366
			Furniture Parts: 3708
			Furniture Accessories: 3712
			Home Furniture: 150303
			Office Furniture: 150304
			Commercial Furniture: 150301
			Outdoor Furniture: 150302
			Children Furniture: 100001203
			Bar Furniture: 100001146
			Beads and Jewelry Making: 200154003
			Jewelry making: 205871206
			Customized Jewelry: 205952104
			Fine Jewelry: 200188001
			Wedding and Engagement Jewelry: 200000161
			Earrings: 200000139
			Necklaces and Pendants: 200000109
			Bracelets and Bangles: 200000097
			Jewelry Sets and More: 200132001
			Rings: 100006749
			Women’s Bracelet Watches: 200214074
			Lover’s Watches: 200214047
			Children’s Watches: 200214043
			Women’s Watches: 200214036
			Men’s Watches: 200214006
			Pocket and Fob Watches: 361120
			Watch Accessories: 200000084
			Human Hair Weaves: 200218141
			Synthetic Extensions: 200217672
			Hair Braids: 200217671
			Synthetic Wigs: 200217666
			Salon Hair Supply Chain: 200217696
			Hair Extensions: 200217614
			DIY Wigs: 201964065
			Hair Salon Tools and Accessories: 200002956
			Lace Wigs: 200004346
			Hair Pieces: 200004940
			Prepaid Digital Code: 205962702
			Tickets: 205966601
			Coupons: 205966201
			Software and Games: 205964202
			Work Wear and Uniforms: 200001355
			Exotic Apparel: 200001271
			Costumes and Accessories: 200001270
			Traditional and Cultural Wear: 200001096
			Stage and Dance Wear: 100003240
			Dresses under $80: 201932007
			Bridesmaid Dresses: 100003270
			Mother of the Bride Dresses: 100005823
			Quinceanera Dresses: 200001556
			Homecoming Dresses: 200001554
			Celebrity-Inspired Dresses: 200001553
			Wedding Party Dress: 200001520
			Wedding Dresses: 100003269
			Cocktail Dresses: 100005790
			Evening Dresses: 100005792
			Prom Dresses: 100005791
			Wedding Accessories: 100005624
			Rompers: 200215341
			Bodysuits: 200215336
			Pants and Capris: 205927403
			Dress: 205871601
			Skirt: 205876401
			Jeans: 100003086,205874801
			Muslim Fashion: 206081401
			Bottoms: 200118010
			Plus size clothes: 206083901
			Swimsuit: 205895301
			Women Tops: 205900902
			Blouses and Shirts: 200001648
			Dresses: 200003482
			Jumpsuits: 200001092
			Sweaters: 200000701,200000783
			Suits and Sets: 200000782
			Jackets and Coats: 200000775,200000662
			Tops and Tees: 200000707,200000785
			Hoodies and Sweatshirts: 100003084,100003141
			Men’s Sets: 200216733
			Pants: 200118008
			Board Shorts: 200000709
			Suits and Blazers: 200000692
			Shirts: 200000668
			Casual Shorts: 100003088
			Women’s Hair Accessories: 200001996
			Women’s Belts: 200003450
			Women’s Hats: 200000769
			Women’s Scarves: 200000743
			Women’s Glasses: 200000741
			Women’s Gloves: 200000732
			Women’s Accessories: 200000724
			Men’s Glasses: 200000617
			Men’s Gloves: 200000609
			Men’s Scarves: 200000613
			Men’s Hats: 200000625
			Men’s Ties and Handkerchiefs: 200000601
			Men’s Belts: 200000600
			Men’s Accessories: 200000599
			Girl’s Accessories: 205778115
			Garment Fabrics and Accessories: 205842201
			Boy’s Accessories: 205780114
			Men’s Socks: 200003491
			Women’s Sleepwears: 200000777
			Women’s Socks and Hosiery: 200000781
			Women’s Intimates: 200000773
			Men’s Underwear: 200000708
			Men’s Sleep and Lounge: 200000673
			Women’s panties: 205780509
			Mobile Phone Accessories: 200084017
			Mobile Phone Parts: 200086021
			Phone Bags & Cases: 200216584
			Mobile Phones: 5090301
			Feature Phones: 5090201
			SIM Cards: 5090101
			## FINAL JSON STRUCTURE:
			{
			"keywords": "string",
			"min_sale_price": number|null,
			"max_sale_price": number|null,
			"original_budget": "string|null",
			"delivery_days": number|null,
			"ship_to_country": "ET",
			"target_currency": "USD",
			"target_language": "en",
			"is_etb": boolean,
			"query_class": "string",
			"language": "am" | "en",
			"category_ids": "string"
			}

			USER PROMPT: %s
			`, userPrompt)
}

func comparsionPrompt(productDetails []*domain.Product, deliveryInfo string, lang string, b string) string {
	return fmt.Sprintf(`You are an expert e-commerce product comparison assistant. Analyze and compare %d products thoroughly.

		CRITICAL INSTRUCTIONS:
		1. Return STRICT JSON ONLY, no additional text or commentary
		2. JSON structure must exactly match the format below
		3. If language is Amharic ("am"), **translate both feature keys and descriptive text to Amharic**
		4: UPDATE aiMatchPercentage FEILD OF EVERY PRODUCT DEPEND ON THE COMPARSION (THE CONFIDENCE OF YOU LLM TO SUGGEST AS BEST PRODUCT FORM THE PRODUCTS)

		{
		"products": [
			{
			"product": { "id":"123", "title":"Sample Product", "imageUrl":"https://example.com/image.jpg", "price":{"etb":1000,"usd":20,"fxTimestamp":"2025-09-03T12:00:00Z"}, "productRating":4.5, "sellerScore":95, "deliveryEstimate":"3-5 days", "description":"Sample description", "summaryBullets":["bullet1","bullet2"], "deeplinkUrl":"https://example.com/product/123", "taxRate":0.1, "discount":10, "numberSold":150, "aiMatchPercentage":0 },
			"synthesis": {
				English (en)
				"pros": [
					"Affordable price",
					"High quality materials",
					"Fast delivery"
				],
				"cons": [
					"Limited color options"
				]

				Amharic (am)
				"pros": [
					"ተመጣጣኝ ዋጋ",
					"ከፍተኛ ጥራት ያላቸው ቁሳቁሶች",
					"ፈጣን እና ታማኝ አሰራር"
				],
				"cons": [
					"የቀለም አማራጮች ገደብ ተደርጓል"
				]
				"isBestValue": true,
				"features": {
				// English keys if lang="en"
				"Price & Value": "Cheaper than competitors with excellent value",
				"Quality & Durability": "Solid build with premium materials and long-lasting",
				"Seller Trust": "Highly rated seller with good reputation",
				"Delivery Speed": "Faster than most competitors",
				"Popularity & Demand": "Well-liked with high sales volume",
				"Unique Features": "Supports wireless charging and extra features",

				// Amharic keys if lang="am"
				"ዋጋ እና የዋጋ እኩልነት": "ከተወዳጅ ምርቶች በተመጣጣኝ ዋጋ ይገኛል",
				"ጥራት እና ቆይታ": "ጥሩ እና ረጅም ቆይታ ያለው ጥራት ተሸማች",
				"የሻጭ እርግጠኝነት": "ከፍተኛ ደረጃ ያለው ታማኝ ሻጭ",
				"የአሰራር ፍጥነት": "ከአብዛኛዎቹ ምርቶች ፈጣን የማድረስ ጊዜ",
				"ታዋቂነት እና ጥያቄ": "በከፍተኛ ብዛት የተሸጠ የታወቀ ምርት",
				"ብቸኛ ስለሆነ ባህሪዎች": "ያለ ግድ የሚከናወን ማስተካከያ ያለው እና ተጨማሪ ባህሪዎች ያሉበት"
				}
			}
			}
			/* repeat above block for all %d products */
		],
		"overallComparison": {
			"bestValueProduct": "Sample Product",
			"bestValueLink": "https://example.com/product/123",
			"bestValuePrice": {
			"etb": 1000,
			"usd": 20,
			"fxTimestamp": "2025-09-03T12:00:00Z"
			},
			"keyHighlights": [
			"Most cost-effective option",
			"High-quality materials and fast delivery"
			],
			"summary": "Sample Product offers the best value, balancing price, quality, and speed, while competitors may offer slightly better features but at higher costs."
		}
		}

		FEATURE COMPARISON GUIDELINES:
		- Compare products relative to each other
		- Highlight strengths, weaknesses, trade-offs, and unique selling points
		- Use descriptive, human-readable phrases
		- Tone: analytical, descriptive, detail-oriented, persuasive
		- If lang='am', feature keys AND values must be in Amharic

		SPECIFIC AREAS TO COMPARE:
		1. PRICE: ETB, USD, discounts, tax
		2. QUALITY: product ratings (%v/5) and customer reviews
		3. SELLER: seller scores (/100) and trust indicators
		4. DELIVERY: delivery estimates ("%s")
		5. POPULARITY: number sold (%d units)
		6. FEATURES: summary bullets and distinctive advantages

		RESPONSE LANGUAGE: %s
		- 'am' → provide all text, including feature keys, in Amharic
		- 'en' → provide all text, including feature keys, in English

		PRODUCTS DATA:
		%s`, len(productDetails), len(productDetails), len(productDetails), deliveryInfo, len(productDetails), lang, string(b))

}
