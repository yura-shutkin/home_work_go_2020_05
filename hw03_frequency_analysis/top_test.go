package hw03_frequency_analysis //nolint:golint

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed
var taskWithAsteriskIsCompleted = false

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{"он", "а", "и", "что", "ты", "не", "если", "то", "его", "кристофер", "робин", "в"}
			require.Subset(t, expected, Top10(text))
		} else {
			expected := []string{"он", "и", "а", "что", "ты", "не", "если", "-", "то", "Кристофер"}
			require.ElementsMatch(t, expected, Top10(text))
		}
	})

	t.Run("positive test 15 words", func(t *testing.T) {
		//words := map[string]int{
		//	"word":       10,
		//	"test":       1,
		//	"Слово":      3,
		//	"слово":      2,
		//	"и":          5,
		//	"словоформа": 15,
		//	"commit":     7,
		//	"git":        12,
		//	"топ":        10,
		//	"задача":     2,
		//	"Задача":     1,
		//	"задача,":    4,
		//	"Тест":       8,
		//	"тест":       1,
		//	"тест.":      3,
		//}

		text = `Тест   word 
		git словоформа топ  задача топ git 
		Тест commit word топ задача, топ слово словоформа 
		словоформа словоформа git задача, git топ word Тест 
		git топ word Тест word словоформа Слово задача, и commit git word 
		 и git Слово задача test задача, словоформа тест. git    тест. 
		словоформа commit commit git тест. git словоформа топ и 
			word словоформа word git словоформа словоформа commit и 
		словоформа git commit Слово словоформа word топ word 
		словоформа тест топ Задача Тест Тест commit и топ Тест слово словоформа Тест`

		expected := []string{"словоформа", "git", "word", "топ", "Тест", "commit", "и", "задача,", "тест.", "Слово"}
		require.ElementsMatch(t, expected, Top10(text))
	})

	t.Run("positive test 1 word", func(t *testing.T) {
		//words := map[string]int{
		//	"word": 10,
		//}

		text = `word word word
		word word
		word
		word    word    word    word`

		expected := []string{"word"}
		require.ElementsMatch(t, expected, Top10(text))
	})

	t.Run("positive test 10 words", func(t *testing.T) {
		//words := map[string]int{
		//	"a":       1,
		//	"bc":      2,
		//	"def":     3,
		//	"single":  4,
		//	"double":  5,
		//	"testing": 6,
		//	"test":    7,
		//	"pass":    8,
		//	"break":   9,
		//	"skip":    10,
		//}

		text = `break double def double skip 
		pass pass test single skip pass testing double 
		testing  test testing testing break bc single 
		pass test skip pass skip pass double def break 
		single skip skip break skip break testing break 
		test break skip bc skip break skip single test
		def double break a testing test test pass pass`

		expected := []string{"skip", "break", "pass", "test", "testing", "double", "single", "def", "bc", "a"}
		require.ElementsMatch(t, expected, Top10(text))
	})
}
