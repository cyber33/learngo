// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Print the literals
//
//  1. Print a few integer literals
//
//  2. Print a few float literals
//
//  3. Print true and false bool constants
//
//  4. Print your name using a string literal
//
//  5. Print a non-english sentence using a string literal
//
// ---------------------------------------------------------

import (
	"fmt"
)

func main() {
	// Use fmt.Println()
	fmt.Println(42)
	fmt.Println(314159)
	fmt.Println(-256)
	fmt.Println(3.14159)
	fmt.Println(true)
	fmt.Println(false)
	fmt.Println("Mark Cyber33")
	fmt.Println("Let's try english and convert this sentence using Google's translate.")
	// Arabic not visible but runs and the cursor goes in the right direction when you move over the text i.e. right to left.
	fmt.Println("لنجرب اللغة الإنجليزية ونحول هذه الجملة باستخدام ترجمة Google.")
	// Spanish is good.
	fmt.Println("Probemos inglés y conviertamos esta oración usando el traductor de Google.")
	// Russian
	fmt.Println("Давайте попробуем по-английски и переведем это предложение с помощью переводчика Google.")
        // Chinese Traditional
	fmt.Println("讓我們試試英語並使用谷歌的翻譯來轉換這句話。")
	// Tamil
	fmt.Println("Google இன் மொழிபெயர்ப்பைப் பயன்படுத்தி இந்த வாக்கியத்தை ஆங்கிலத்தில் மாற்ற முயற்சிப்போம்.")
	// Urdu not so much but runs and the cursor goes in the right direction when you move over the text i.e. right to left. 
	fmt.Println("آئیے انگریزی کی کوشش کریں اور گوگل کے ترجمہ کا استعمال کرتے ہوئے اس جملے کو تبدیل کریں۔")
        // Swahili
	fmt.Println("Hebu tujaribu Kiingereza na tubadili sentensi hii kwa kutumia tafsiri ya Google.")
	// Korean
	fmt.Println("Google 번역을 사용하여 영어를 시도하고 이 문장을 변환해 봅시다.")
	// Persian not correct visually but runs and the cursor goes in the right direction when you move over the text i.e. right to left.
	fmt.Println("بیایید انگلیسی را امتحان کنیم و این جمله را با استفاده از ترجمه گوگل تبدیل کنیم.")
	// Kurdish (Kurmanji) 
	fmt.Println("Ka em îngilîzî biceribînin û vê hevokê bi wergera Google-ê veguherînin.")
}
