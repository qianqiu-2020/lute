// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package test

import (
	"testing"

	"github.com/88250/lute"
)

var spinVditorSVDOMTests = []*parseTest{
	//{"9", "* foo\n  * bar\n* b‸", "<span data-tight=\"true\" data-type=\"ul\" data-marker=\"*\" data-block=\"0\"><span data-type=\"li\" data-space=\"\"><span data-type=\"li-marker\" class=\"vditor-sv__marker--bi\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"li\" data-space=\"  \"><span data-type=\"li-space\">  </span><span data-type=\"li-marker\" class=\"vditor-sv__marker--bi\">* </span><span data-type=\"text\">bar</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span><span data-type=\"li\" data-space=\"\"><span data-type=\"li-marker\" class=\"vditor-sv__marker--bi\">* </span><span data-type=\"text\">b<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},

	//{"34", "* foo\n  1. ba‸", "<span data-block=\"0\" data-type=\"blockquote\"><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li\" data-space=\"  \"><span data-type=\"li-marker\" class=\"vditor-sv__marker--bi\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li-space\">  </span><span data-type=\"text\">bar<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"33", "> * foo\n> \n>   bar‸\n", "<span data-block=\"0\" data-type=\"blockquote\"><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li\" data-space=\"  \"><span data-type=\"li-marker\" class=\"vditor-sv__marker--bi\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li-space\">  </span><span data-type=\"text\">bar<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"32", "> * foo\n> * bar‸", "<span data-block=\"0\" data-type=\"blockquote\"><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li\" data-space=\"  \"><span data-type=\"li-marker\" class=\"vditor-sv__marker--bi\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span></span><span data-type=\"li\" data-space=\"  \"><span data-type=\"li-marker\" class=\"vditor-sv__marker--bi\">* </span><span data-type=\"text\">bar<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"31", "    f‸", "<span data-block=\"0\" data-type=\"code-block\"><span data-type=\"code-block-open-marker\" class=\"vditor-sv__marker\">```</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span>f<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span class=\"vditor-sv__marker--info\" data-type=\"code-block-info\">```</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"30", "```  `code`  ```‸", "<span data-type=\"p\" data-block=\"0\"><span data-type=\"code\"><span class=\"vditor-sv__marker\">```</span><span> `code` </span><span class=\"vditor-sv__marker\">```</span></span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"29", "`` `code` ``‸", "<span data-type=\"p\" data-block=\"0\"><span data-type=\"code\"><span class=\"vditor-sv__marker\">`` </span><span>`code`</span><span class=\"vditor-sv__marker\"> ``</span></span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"28", "# foo {id}b‸", "<span data-block=\"0\" data-type=\"heading\"><span class=\"h1\"><span class=\"vditor-sv__marker--heading\" data-type=\"heading-marker\"># </span><span data-type=\"text\">foo {id}b<wbr></span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"27", "* foo\n  ```‸", "<span data-tight=\"true\" data-type=\"ul\" data-marker=\"*\" data-block=\"0\"><span data-type=\"li\" data-space=\"\"><span data-type=\"li-marker\" class=\"vditor-sv__marker--bi\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"li-space\">  </span><span data-block=\"0\" data-type=\"code-block\"><span data-type=\"code-block-open-marker\" class=\"vditor-sv__marker\">```</span><span class=\"vditor-sv__marker--info\" data-type=\"code-block-info\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"li-space\">  </span><span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"li-space\">  </span><span data-type=\"code-block-close-marker\" class=\"vditor-sv__marker\">```</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"26", "> foo\n> \n> > bar\n> \n> baz‸", "<span data-block=\"0\" data-type=\"blockquote\"><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">bar</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">baz<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"25", "foo\n> bar‸", "<span data-type=\"p\" data-block=\"0\"><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span><span data-block=\"0\" data-type=\"blockquote\"><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">bar<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"24", "# foo {id‸}", "<span data-block=\"0\" data-type=\"heading\"><span class=\"h1\"><span class=\"vditor-sv__marker--heading\" data-type=\"heading-marker\"># </span><span data-type=\"text\">foo</span><span class=\"vditor-sv__marker\"> {id<wbr>}</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"23", "# foo {id}‸", "<span data-block=\"0\" data-type=\"heading\"><span class=\"h1\"><span class=\"vditor-sv__marker--heading\" data-type=\"heading-marker\"># </span><span data-type=\"text\">foo</span><span class=\"vditor-sv__marker\"> {id}</span><span data-type=\"text\"><wbr></span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"22", "```\nfoo\n```‸", "<span data-block=\"0\" data-type=\"code-block\"><span data-type=\"code-block-open-marker\" class=\"vditor-sv__marker\">```</span><span class=\"vditor-sv__marker--info\" data-type=\"code-block-info\"></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span>foo<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"code-block-close-marker\" class=\"vditor-sv__marker\">```</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"21", "> foo\n> >‸\n", "<span data-block=\"0\" data-type=\"blockquote\"><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"20", "> >‸", "<span data-block=\"0\" data-type=\"blockquote\"><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"19", "这里是一个脚注引用[^1]，这里是另一个脚注引用[^bignote]。\n\n[^1]: 第一个脚注定义。\n[^bignote]: 脚注定义可使用多段内容。\n\n    缩进对齐的段落包含在这个脚注定义内。‸\n", "<span data-type=\"p\" data-block=\"0\"><span data-type=\"text\">这里是一个脚注引用</span><span class=\"sup\"><span class=\"vditor-sv__marker--bracket\">[</span><span class=\"vditor-sv__marker--link\">^1</span><span class=\"vditor-sv__marker--bracket\">]</span></span><span data-type=\"text\">，这里是另一个脚注引用</span><span class=\"sup\"><span class=\"vditor-sv__marker--bracket\">[</span><span class=\"vditor-sv__marker--link\">^bignote</span><span class=\"vditor-sv__marker--bracket\">]</span></span><span data-type=\"text\">。</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span><span data-block=\"0\" data-type=\"footnotes-block\"><span data-type=\"footnotes-def\"><span class=\"vditor-sv__marker--bracket\">[</span><span class=\"vditor-sv__marker--link\" data-type=\"footnotes-link\">^1</span><span class=\"vditor-sv__marker--bracket\">]</span>: <span data-type=\"text\">第一个脚注定义。</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span><span data-type=\"footnotes-def\"><span class=\"vditor-sv__marker--bracket\">[</span><span class=\"vditor-sv__marker--link\" data-type=\"footnotes-link\">^bignote</span><span class=\"vditor-sv__marker--bracket\">]</span>: <span data-type=\"text\">脚注定义可使用多段内容。</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"footnotes-space\">    </span><span data-type=\"text\">缩进对齐的段落包含在这个脚注定义内。<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span></span>"},
	//{"18", "|foo|bar|\n| ---| ---|\n|‸", "<span data-block=\"0\" data-type=\"table\">|foo|bar|\n| ---| ---|\n|<wbr><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"17", "‸", "<span data-type=\"p\" data-block=\"0\"><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"16", "> * f‸\n> \n", "<span data-block=\"0\" data-type=\"blockquote\"><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li\" data-space=\"  \"><span data-type=\"li-marker\" class=\"vditor-sv__marker--bi\">* </span><span data-type=\"text\">f<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"15", "> ## foo\n>  \n> * b‸", "<span data-block=\"0\" data-type=\"blockquote\"><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span class=\"vditor-sv__marker--heading\" data-type=\"heading-marker\">## </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li\" data-space=\"  \"><span data-type=\"li-marker\" class=\"vditor-sv__marker--bi\">* </span><span data-type=\"text\">b<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"14", "> foo\n> \n> b‸", "<span data-block=\"0\" data-type=\"blockquote\"><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">b<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"13", "> * foo\n>‸", "<span data-block=\"0\" data-type=\"blockquote\"><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li\" data-space=\"  \"><span data-type=\"li-marker\" class=\"vditor-sv__marker--bi\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"12", "> * f‸", "<span data-block=\"0\" data-type=\"blockquote\"><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li\" data-space=\"  \"><span data-type=\"li-marker\" class=\"vditor-sv__marker--bi\">* </span><span data-type=\"text\">f<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	//{"10", "* foo\n\n  bar\n\n* baz\n\n  b‸", "<span data-type=\"ul\" data-marker=\"*\" data-block=\"0\"><span data-type=\"li\" data-space=\"\"><span data-type=\"li-marker\" class=\"vditor-sv__marker--bi\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"li-space\">  </span><span data-type=\"text\">bar</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span><span data-type=\"li\" data-space=\"\"><span data-type=\"li-marker\" class=\"vditor-sv__marker--bi\">* </span><span data-type=\"text\">baz</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"li-space\">  </span><span data-type=\"text\">b<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	{"9", "* foo\n\n  bar‸", "<span data-type=\"ul\" data-marker=\"*\" data-block=\"0\"><span data-type=\"li\" data-space=\"\"><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"p\" data-block=\"0\"><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"li-space\">  </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"li-space\">  </span></span><span data-type=\"p\" data-block=\"0\"><span data-type=\"text\">bar<wbr></span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	{"8", "* foo\n\n‸\n", "<span data-tight=\"true\" data-type=\"ul\" data-marker=\"*\" data-block=\"0\"><span data-type=\"li\" data-space=\"\"><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"p\" data-block=\"0\"><span data-type=\"text\">foo</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span><span data-type=\"p\" data-block=\"0\"><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	{"7", "*‸", "<span data-type=\"p\" data-block=\"0\"><span data-type=\"text\">*<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	{"6", "> * foo\n> \n>‸", "<span data-block=\"0\" data-type=\"blockquote\"><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-tight=\"true\" data-type=\"ul\" data-marker=\"*\" data-block=\"0\"><span data-type=\"li\" data-space=\"  \"><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"p\" data-block=\"0\"><span data-type=\"text\">foo</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span></span><span data-type=\"p\" data-block=\"0\"><span data-type=\"text\"><wbr></span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	{"5", "> foo\n> # b‸", "<span data-block=\"0\" data-type=\"blockquote\"><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"p\" data-block=\"0\"><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span></span><span class=\"vditor-sv__marker--heading\" data-type=\"heading-marker\"># </span><span data-type=\"text\">b<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	{"4", "* foo\n  * ba‸", "<span data-tight=\"true\" data-type=\"ul\" data-marker=\"*\" data-block=\"0\"><span data-type=\"li\" data-space=\"\"><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"p\" data-block=\"0\"><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"li-space\">  </span></span><span data-tight=\"true\" data-type=\"ul\" data-marker=\"*\" data-block=\"0\"><span data-type=\"li\" data-space=\"  \"><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"p\" data-block=\"0\"><span data-type=\"text\">ba<wbr></span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"li-space\">  </span></span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	{"3", "* foo\n‸", "<span data-tight=\"true\" data-type=\"ul\" data-marker=\"*\" data-block=\"0\"><span data-type=\"li\" data-space=\"\"><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"p\" data-block=\"0\"><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"li-space\">  </span><span data-type=\"text\"><wbr></span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	{"2", "> foo\n> \n> bar‸", "<span data-block=\"0\" data-type=\"blockquote\"><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"p\" data-block=\"0\"><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span></span><span data-type=\"p\" data-block=\"0\"><span data-type=\"text\">bar<wbr></span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	{"1", "> foo‸", "<span data-block=\"0\" data-type=\"blockquote\"><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"p\" data-block=\"0\"><span data-type=\"text\">foo<wbr></span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	{"0", "*foo*‸", "<span data-type=\"p\" data-block=\"0\"><span data-type=\"em\" class=\"em\"><span class=\"vditor-sv__marker--bi\">*</span><span data-type=\"text\">foo</span><span class=\"vditor-sv__marker--bi\">*</span></span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
}

func TestSpinVditorSVDOM(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.ToC = true

	for _, test := range spinVditorSVDOMTests {
		html := luteEngine.SpinVditorSVDOM(test.from)
		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal html\n\t%q", test.name, test.to, html, test.from)
		}
	}
}
