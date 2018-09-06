//line yuri.rl:1
// Copyright (c) 2018, Hayden Eskriett <hayden@eskriett.com>
// See LICENSE for licensing details
//
// Build with:
//   ragel -G2 -Z yuri.rl

// Package yuri provides a blazing fast way to extracts URIs from plain text
package yuri

// YankURIs extracts URIs from text
func YankURIs(data []byte) []string {

	cs, p, pe, eof := 0, 0, len(data), len(data)
	ts, te, act := 0, 0, 0
	_ = eof
	_ = ts
	_ = te
	_ = act

	var uris []string

//line yuri.go:26
	const yuriStart int = 179
//line yuri.rl:84

//line yuri.go:38
	{
		cs = yuriStart
		ts = 0
		te = 0
		act = 0
	}

//line yuri.rl:87

//line yuri.go:48
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 179:
			goto st_case_179
		case 180:
			goto st_case_180
		case 0:
			goto st_case_0
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 181:
			goto st_case_181
		case 9:
			goto st_case_9
		case 182:
			goto st_case_182
		case 183:
			goto st_case_183
		case 184:
			goto st_case_184
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 185:
			goto st_case_185
		case 186:
			goto st_case_186
		case 187:
			goto st_case_187
		case 188:
			goto st_case_188
		case 189:
			goto st_case_189
		case 190:
			goto st_case_190
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 191:
			goto st_case_191
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 65:
			goto st_case_65
		case 66:
			goto st_case_66
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 69:
			goto st_case_69
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		case 73:
			goto st_case_73
		case 74:
			goto st_case_74
		case 75:
			goto st_case_75
		case 76:
			goto st_case_76
		case 77:
			goto st_case_77
		case 78:
			goto st_case_78
		case 79:
			goto st_case_79
		case 80:
			goto st_case_80
		case 81:
			goto st_case_81
		case 82:
			goto st_case_82
		case 83:
			goto st_case_83
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
		case 86:
			goto st_case_86
		case 87:
			goto st_case_87
		case 88:
			goto st_case_88
		case 89:
			goto st_case_89
		case 90:
			goto st_case_90
		case 91:
			goto st_case_91
		case 92:
			goto st_case_92
		case 93:
			goto st_case_93
		case 94:
			goto st_case_94
		case 95:
			goto st_case_95
		case 96:
			goto st_case_96
		case 97:
			goto st_case_97
		case 98:
			goto st_case_98
		case 99:
			goto st_case_99
		case 100:
			goto st_case_100
		case 101:
			goto st_case_101
		case 102:
			goto st_case_102
		case 103:
			goto st_case_103
		case 104:
			goto st_case_104
		case 105:
			goto st_case_105
		case 106:
			goto st_case_106
		case 107:
			goto st_case_107
		case 108:
			goto st_case_108
		case 109:
			goto st_case_109
		case 110:
			goto st_case_110
		case 111:
			goto st_case_111
		case 112:
			goto st_case_112
		case 113:
			goto st_case_113
		case 114:
			goto st_case_114
		case 115:
			goto st_case_115
		case 116:
			goto st_case_116
		case 117:
			goto st_case_117
		case 118:
			goto st_case_118
		case 119:
			goto st_case_119
		case 120:
			goto st_case_120
		case 121:
			goto st_case_121
		case 122:
			goto st_case_122
		case 123:
			goto st_case_123
		case 124:
			goto st_case_124
		case 125:
			goto st_case_125
		case 126:
			goto st_case_126
		case 127:
			goto st_case_127
		case 128:
			goto st_case_128
		case 129:
			goto st_case_129
		case 130:
			goto st_case_130
		case 131:
			goto st_case_131
		case 132:
			goto st_case_132
		case 133:
			goto st_case_133
		case 134:
			goto st_case_134
		case 135:
			goto st_case_135
		case 136:
			goto st_case_136
		case 137:
			goto st_case_137
		case 138:
			goto st_case_138
		case 139:
			goto st_case_139
		case 140:
			goto st_case_140
		case 141:
			goto st_case_141
		case 142:
			goto st_case_142
		case 143:
			goto st_case_143
		case 144:
			goto st_case_144
		case 145:
			goto st_case_145
		case 146:
			goto st_case_146
		case 147:
			goto st_case_147
		case 148:
			goto st_case_148
		case 149:
			goto st_case_149
		case 150:
			goto st_case_150
		case 151:
			goto st_case_151
		case 152:
			goto st_case_152
		case 153:
			goto st_case_153
		case 154:
			goto st_case_154
		case 155:
			goto st_case_155
		case 156:
			goto st_case_156
		case 157:
			goto st_case_157
		case 158:
			goto st_case_158
		case 159:
			goto st_case_159
		case 160:
			goto st_case_160
		case 161:
			goto st_case_161
		case 162:
			goto st_case_162
		case 163:
			goto st_case_163
		case 164:
			goto st_case_164
		case 165:
			goto st_case_165
		case 166:
			goto st_case_166
		case 167:
			goto st_case_167
		case 168:
			goto st_case_168
		case 192:
			goto st_case_192
		case 169:
			goto st_case_169
		case 193:
			goto st_case_193
		case 194:
			goto st_case_194
		case 195:
			goto st_case_195
		case 196:
			goto st_case_196
		case 197:
			goto st_case_197
		case 198:
			goto st_case_198
		case 199:
			goto st_case_199
		case 200:
			goto st_case_200
		case 170:
			goto st_case_170
		case 171:
			goto st_case_171
		case 172:
			goto st_case_172
		case 173:
			goto st_case_173
		case 201:
			goto st_case_201
		case 174:
			goto st_case_174
		case 175:
			goto st_case_175
		case 176:
			goto st_case_176
		case 177:
			goto st_case_177
		case 178:
			goto st_case_178
		}
		goto st_out
	tr0:
//line yuri.rl:82
		p = (te) - 1

		goto st179
	tr10:
//line NONE:1
		switch act {
		case 1:
			{
				p = (te) - 1
				uris = append(uris, string(data[ts:te]))
			}
		default:
			{
				p = (te) - 1
			}
		}

		goto st179
	tr13:
//line yuri.rl:26
		p = (te) - 1
		{
			uris = append(uris, string(data[ts:te]))
		}
		goto st179
	tr183:
//line yuri.rl:82
		te = p + 1

		goto st179
	tr187:
//line yuri.rl:82
		te = p
		p--

		goto st179
	tr189:
//line yuri.rl:26
		te = p
		p--
		{
			uris = append(uris, string(data[ts:te]))
		}
		goto st179
	st179:
//line NONE:1
		ts = 0

		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
//line NONE:1
		ts = p

//line yuri.go:510
		switch data[p] {
		case 70:
			goto tr184
		case 72:
			goto tr185
		case 77:
			goto tr186
		case 102:
			goto tr184
		case 104:
			goto tr185
		case 109:
			goto tr186
		}
		goto tr183
	tr184:
//line NONE:1
		te = p + 1

//line yuri.rl:82
		act = 2
		goto st180
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
//line yuri.go:538
		switch data[p] {
		case 84:
			goto st0
		case 116:
			goto st0
		}
		goto tr187
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 80:
			goto st1
		case 112:
			goto st1
		}
		goto tr0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 58 {
			goto st2
		}
		goto tr0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 47 {
			goto st3
		}
		goto tr0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 47 {
			goto st4
		}
		goto tr0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 33:
			goto st5
		case 37:
			goto st6
		case 47:
			goto tr0
		case 60:
			goto tr0
		case 64:
			goto st8
		case 91:
			goto st14
		case 95:
			goto st5
		case 126:
			goto st5
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 35:
				if 36 <= data[p] && data[p] <= 46 {
					goto st5
				}
			default:
				goto tr0
			}
		case data[p] > 61:
			switch {
			case data[p] < 92:
				if 62 <= data[p] && data[p] <= 63 {
					goto tr0
				}
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr0
				}
			default:
				goto tr0
			}
		default:
			goto st5
		}
		goto tr7
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 37:
			goto st6
		case 47:
			goto tr10
		case 60:
			goto tr10
		case 64:
			goto st8
		case 96:
			goto tr10
		case 127:
			goto tr10
		}
		switch {
		case data[p] < 62:
			switch {
			case data[p] > 32:
				if 34 <= data[p] && data[p] <= 35 {
					goto tr10
				}
			default:
				goto tr10
			}
		case data[p] > 63:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 125 {
					goto tr10
				}
			case data[p] >= 91:
				goto tr10
			}
		default:
			goto tr10
		}
		goto st5
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st7
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st7
			}
		default:
			goto st7
		}
		goto tr10
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st5
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st5
			}
		default:
			goto st5
		}
		goto tr10
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 91 {
			goto st14
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr10
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr10
				}
			case data[p] >= 92:
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr12
	tr12:
//line NONE:1
		te = p + 1

		goto st181
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
//line yuri.go:749
		switch data[p] {
		case 47:
			goto tr19
		case 58:
			goto st185
		case 96:
			goto tr189
		case 127:
			goto tr189
		}
		switch {
		case data[p] < 59:
			if data[p] <= 44 {
				goto tr189
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 125 {
					goto tr189
				}
			case data[p] >= 91:
				goto tr189
			}
		default:
			goto tr189
		}
		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 47:
			goto tr13
		case 95:
			goto st9
		case 126:
			goto st9
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 44:
				if 45 <= data[p] && data[p] <= 46 {
					goto st9
				}
			default:
				goto tr13
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr13
				}
			case data[p] >= 91:
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr15
	tr15:
//line NONE:1
		te = p + 1

		goto st182
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
//line yuri.go:824
		switch data[p] {
		case 47:
			goto tr19
		case 58:
			goto st185
		case 95:
			goto st9
		case 126:
			goto st9
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 44:
				if 45 <= data[p] && data[p] <= 46 {
					goto st9
				}
			default:
				goto tr189
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr189
				}
			case data[p] >= 91:
				goto tr189
			}
		default:
			goto tr189
		}
		goto tr15
	tr19:
//line NONE:1
		te = p + 1

		goto st183
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
//line yuri.go:868
		switch data[p] {
		case 34:
			goto tr189
		case 35:
			goto tr17
		case 37:
			goto st12
		case 60:
			goto tr189
		case 62:
			goto tr189
		case 96:
			goto tr189
		case 127:
			goto tr189
		}
		switch {
		case data[p] < 91:
			if data[p] <= 32 {
				goto tr189
			}
		case data[p] > 94:
			if 123 <= data[p] && data[p] <= 125 {
				goto tr189
			}
		default:
			goto tr189
		}
		goto tr19
	tr17:
//line NONE:1
		te = p + 1

		goto st184
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
//line yuri.go:908
		switch data[p] {
		case 37:
			goto st10
		case 60:
			goto tr189
		case 62:
			goto tr189
		case 96:
			goto tr189
		case 127:
			goto tr189
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto tr189
			}
		case data[p] > 35:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 125 {
					goto tr189
				}
			case data[p] >= 91:
				goto tr189
			}
		default:
			goto tr189
		}
		goto tr17
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr13
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr17
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr13
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st13
			}
		default:
			goto st13
		}
		goto tr13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr19
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr13
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
		if data[p] == 47 {
			goto tr19
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st186
		}
		goto tr189
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		if data[p] == 47 {
			goto tr19
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st187
		}
		goto tr189
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
		if data[p] == 47 {
			goto tr19
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st188
		}
		goto tr189
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
		if data[p] == 47 {
			goto tr19
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st189
		}
		goto tr189
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
		if data[p] == 47 {
			goto tr19
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st190
		}
		goto tr189
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
		if data[p] == 47 {
			goto tr19
		}
		goto tr189
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 58:
			goto st150
		case 118:
			goto st165
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st15
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st15
			}
		default:
			goto st15
		}
		goto tr10
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if data[p] == 58 {
			goto st19
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st16
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr10
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if data[p] == 58 {
			goto st19
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st17
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st17
			}
		default:
			goto st17
		}
		goto tr10
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if data[p] == 58 {
			goto st19
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st18
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st18
			}
		default:
			goto st18
		}
		goto tr10
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if data[p] == 58 {
			goto st19
		}
		goto tr10
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 58 {
			goto st136
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st20
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st20
			}
		default:
			goto st20
		}
		goto tr10
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 58 {
			goto st24
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st21
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st21
			}
		default:
			goto st21
		}
		goto tr10
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if data[p] == 58 {
			goto st24
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st22
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st22
			}
		default:
			goto st22
		}
		goto tr10
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		if data[p] == 58 {
			goto st24
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st23
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st23
			}
		default:
			goto st23
		}
		goto tr10
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if data[p] == 58 {
			goto st24
		}
		goto tr10
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 58 {
			goto st122
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st25
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st25
			}
		default:
			goto st25
		}
		goto tr10
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if data[p] == 58 {
			goto st29
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr10
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if data[p] == 58 {
			goto st29
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st27
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st27
			}
		default:
			goto st27
		}
		goto tr10
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if data[p] == 58 {
			goto st29
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st28
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st28
			}
		default:
			goto st28
		}
		goto tr10
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		if data[p] == 58 {
			goto st29
		}
		goto tr10
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		if data[p] == 58 {
			goto st108
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st30
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st30
			}
		default:
			goto st30
		}
		goto tr10
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		if data[p] == 58 {
			goto st34
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st31
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st31
			}
		default:
			goto st31
		}
		goto tr10
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		if data[p] == 58 {
			goto st34
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st32
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st32
			}
		default:
			goto st32
		}
		goto tr10
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if data[p] == 58 {
			goto st34
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st33
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st33
			}
		default:
			goto st33
		}
		goto tr10
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		if data[p] == 58 {
			goto st34
		}
		goto tr10
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		if data[p] == 58 {
			goto st94
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st35
			}
		default:
			goto st35
		}
		goto tr10
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		if data[p] == 58 {
			goto st39
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st36
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st36
			}
		default:
			goto st36
		}
		goto tr10
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		if data[p] == 58 {
			goto st39
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st37
			}
		default:
			goto st37
		}
		goto tr10
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		if data[p] == 58 {
			goto st39
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st38
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st38
			}
		default:
			goto st38
		}
		goto tr10
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		if data[p] == 58 {
			goto st39
		}
		goto tr10
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		if data[p] == 58 {
			goto st80
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st40
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st40
			}
		default:
			goto st40
		}
		goto tr10
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		if data[p] == 58 {
			goto st44
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st41
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st41
			}
		default:
			goto st41
		}
		goto tr10
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		if data[p] == 58 {
			goto st44
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st42
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st42
			}
		default:
			goto st42
		}
		goto tr10
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if data[p] == 58 {
			goto st44
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st43
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st43
			}
		default:
			goto st43
		}
		goto tr10
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		if data[p] == 58 {
			goto st44
		}
		goto tr10
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch data[p] {
		case 48:
			goto st45
		case 49:
			goto st71
		case 50:
			goto st74
		case 58:
			goto st78
		}
		switch {
		case data[p] < 65:
			if 51 <= data[p] && data[p] <= 57 {
				goto st77
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st79
			}
		default:
			goto st79
		}
		goto tr10
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st67
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st64
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st64
			}
		default:
			goto st64
		}
		goto tr10
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 48:
			goto st47
		case 49:
			goto st60
		case 50:
			goto st62
		}
		if 51 <= data[p] && data[p] <= 57 {
			goto st61
		}
		goto tr10
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		if data[p] == 46 {
			goto st48
		}
		goto tr10
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 48:
			goto st49
		case 49:
			goto st56
		case 50:
			goto st58
		}
		if 51 <= data[p] && data[p] <= 57 {
			goto st57
		}
		goto tr10
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		if data[p] == 46 {
			goto st50
		}
		goto tr10
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 48:
			goto st51
		case 49:
			goto st52
		case 50:
			goto st54
		}
		if 51 <= data[p] && data[p] <= 57 {
			goto st53
		}
		goto tr10
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		if data[p] == 93 {
			goto st191
		}
		goto tr10
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
		switch data[p] {
		case 47:
			goto tr19
		case 58:
			goto st185
		}
		goto tr189
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		if data[p] == 93 {
			goto st191
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st53
		}
		goto tr10
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		if data[p] == 93 {
			goto st191
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st51
		}
		goto tr10
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 53:
			goto st55
		case 93:
			goto st191
		}
		switch {
		case data[p] > 52:
			if 54 <= data[p] && data[p] <= 57 {
				goto st51
			}
		case data[p] >= 48:
			goto st53
		}
		goto tr10
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		if data[p] == 93 {
			goto st191
		}
		if 48 <= data[p] && data[p] <= 53 {
			goto st51
		}
		goto tr10
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		if data[p] == 46 {
			goto st50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st57
		}
		goto tr10
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		if data[p] == 46 {
			goto st50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st49
		}
		goto tr10
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 46:
			goto st50
		case 53:
			goto st59
		}
		switch {
		case data[p] > 52:
			if 54 <= data[p] && data[p] <= 57 {
				goto st49
			}
		case data[p] >= 48:
			goto st57
		}
		goto tr10
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		if data[p] == 46 {
			goto st50
		}
		if 48 <= data[p] && data[p] <= 53 {
			goto st49
		}
		goto tr10
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		if data[p] == 46 {
			goto st48
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st61
		}
		goto tr10
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		if data[p] == 46 {
			goto st48
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st47
		}
		goto tr10
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 46:
			goto st48
		case 53:
			goto st63
		}
		switch {
		case data[p] > 52:
			if 54 <= data[p] && data[p] <= 57 {
				goto st47
			}
		case data[p] >= 48:
			goto st61
		}
		goto tr10
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		if data[p] == 46 {
			goto st48
		}
		if 48 <= data[p] && data[p] <= 53 {
			goto st47
		}
		goto tr10
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		if data[p] == 58 {
			goto st67
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st65
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st65
			}
		default:
			goto st65
		}
		goto tr10
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		if data[p] == 58 {
			goto st67
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st66
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st66
			}
		default:
			goto st66
		}
		goto tr10
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		if data[p] == 58 {
			goto st67
		}
		goto tr10
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		if data[p] == 58 {
			goto st51
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st68
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st68
			}
		default:
			goto st68
		}
		goto tr10
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		if data[p] == 93 {
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st69
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st69
			}
		default:
			goto st69
		}
		goto tr10
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		if data[p] == 93 {
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st70
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st70
			}
		default:
			goto st70
		}
		goto tr10
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		if data[p] == 93 {
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st51
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st51
			}
		default:
			goto st51
		}
		goto tr10
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st67
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st72
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st64
			}
		default:
			goto st64
		}
		goto tr10
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st67
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st73
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st65
			}
		default:
			goto st65
		}
		goto tr10
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st67
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st66
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st66
			}
		default:
			goto st66
		}
		goto tr10
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch data[p] {
		case 46:
			goto st46
		case 53:
			goto st75
		case 58:
			goto st67
		}
		switch {
		case data[p] < 54:
			if 48 <= data[p] && data[p] <= 52 {
				goto st72
			}
		case data[p] > 57:
			switch {
			case data[p] > 70:
				if 97 <= data[p] && data[p] <= 102 {
					goto st64
				}
			case data[p] >= 65:
				goto st64
			}
		default:
			goto st76
		}
		goto tr10
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st67
		}
		switch {
		case data[p] < 54:
			if 48 <= data[p] && data[p] <= 53 {
				goto st73
			}
		case data[p] > 57:
			switch {
			case data[p] > 70:
				if 97 <= data[p] && data[p] <= 102 {
					goto st65
				}
			case data[p] >= 65:
				goto st65
			}
		default:
			goto st65
		}
		goto tr10
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st67
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st65
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st65
			}
		default:
			goto st65
		}
		goto tr10
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st67
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st76
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st64
			}
		default:
			goto st64
		}
		goto tr10
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		if data[p] == 93 {
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st68
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st68
			}
		default:
			goto st68
		}
		goto tr10
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		if data[p] == 58 {
			goto st67
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st64
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st64
			}
		default:
			goto st64
		}
		goto tr10
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		switch data[p] {
		case 48:
			goto st81
		case 49:
			goto st86
		case 50:
			goto st89
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 51 <= data[p] && data[p] <= 57 {
				goto st92
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st93
			}
		default:
			goto st93
		}
		goto tr10
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st85
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st82
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st82
			}
		default:
			goto st82
		}
		goto tr10
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		switch data[p] {
		case 58:
			goto st85
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st83
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st83
			}
		default:
			goto st83
		}
		goto tr10
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 58:
			goto st85
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st84
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st84
			}
		default:
			goto st84
		}
		goto tr10
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		switch data[p] {
		case 58:
			goto st85
		case 93:
			goto st191
		}
		goto tr10
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st68
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st68
			}
		default:
			goto st68
		}
		goto tr10
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st85
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st87
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st82
			}
		default:
			goto st82
		}
		goto tr10
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st85
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st88
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st83
			}
		default:
			goto st83
		}
		goto tr10
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st85
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st84
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st84
			}
		default:
			goto st84
		}
		goto tr10
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 46:
			goto st46
		case 53:
			goto st90
		case 58:
			goto st85
		case 93:
			goto st191
		}
		switch {
		case data[p] < 54:
			if 48 <= data[p] && data[p] <= 52 {
				goto st87
			}
		case data[p] > 57:
			switch {
			case data[p] > 70:
				if 97 <= data[p] && data[p] <= 102 {
					goto st82
				}
			case data[p] >= 65:
				goto st82
			}
		default:
			goto st91
		}
		goto tr10
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st85
		case 93:
			goto st191
		}
		switch {
		case data[p] < 54:
			if 48 <= data[p] && data[p] <= 53 {
				goto st88
			}
		case data[p] > 57:
			switch {
			case data[p] > 70:
				if 97 <= data[p] && data[p] <= 102 {
					goto st83
				}
			case data[p] >= 65:
				goto st83
			}
		default:
			goto st83
		}
		goto tr10
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st85
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st83
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st83
			}
		default:
			goto st83
		}
		goto tr10
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st85
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st91
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st82
			}
		default:
			goto st82
		}
		goto tr10
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 58:
			goto st85
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st82
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st82
			}
		default:
			goto st82
		}
		goto tr10
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch data[p] {
		case 48:
			goto st95
		case 49:
			goto st100
		case 50:
			goto st103
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 51 <= data[p] && data[p] <= 57 {
				goto st106
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st107
			}
		default:
			goto st107
		}
		goto tr10
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st99
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st96
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st96
			}
		default:
			goto st96
		}
		goto tr10
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 58:
			goto st99
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st97
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st97
			}
		default:
			goto st97
		}
		goto tr10
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		switch data[p] {
		case 58:
			goto st99
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st98
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st98
			}
		default:
			goto st98
		}
		goto tr10
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch data[p] {
		case 58:
			goto st99
		case 93:
			goto st191
		}
		goto tr10
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		switch data[p] {
		case 48:
			goto st81
		case 49:
			goto st86
		case 50:
			goto st89
		}
		switch {
		case data[p] < 65:
			if 51 <= data[p] && data[p] <= 57 {
				goto st92
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st93
			}
		default:
			goto st93
		}
		goto tr10
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st99
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st96
			}
		default:
			goto st96
		}
		goto tr10
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st99
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st102
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st97
			}
		default:
			goto st97
		}
		goto tr10
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st99
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st98
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st98
			}
		default:
			goto st98
		}
		goto tr10
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 46:
			goto st46
		case 53:
			goto st104
		case 58:
			goto st99
		case 93:
			goto st191
		}
		switch {
		case data[p] < 54:
			if 48 <= data[p] && data[p] <= 52 {
				goto st101
			}
		case data[p] > 57:
			switch {
			case data[p] > 70:
				if 97 <= data[p] && data[p] <= 102 {
					goto st96
				}
			case data[p] >= 65:
				goto st96
			}
		default:
			goto st105
		}
		goto tr10
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st99
		case 93:
			goto st191
		}
		switch {
		case data[p] < 54:
			if 48 <= data[p] && data[p] <= 53 {
				goto st102
			}
		case data[p] > 57:
			switch {
			case data[p] > 70:
				if 97 <= data[p] && data[p] <= 102 {
					goto st97
				}
			case data[p] >= 65:
				goto st97
			}
		default:
			goto st97
		}
		goto tr10
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st99
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st97
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st97
			}
		default:
			goto st97
		}
		goto tr10
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st99
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st105
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st96
			}
		default:
			goto st96
		}
		goto tr10
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch data[p] {
		case 58:
			goto st99
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st96
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st96
			}
		default:
			goto st96
		}
		goto tr10
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		switch data[p] {
		case 48:
			goto st109
		case 49:
			goto st114
		case 50:
			goto st117
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 51 <= data[p] && data[p] <= 57 {
				goto st120
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st121
			}
		default:
			goto st121
		}
		goto tr10
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st113
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st110
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st110
			}
		default:
			goto st110
		}
		goto tr10
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 58:
			goto st113
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st111
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st111
			}
		default:
			goto st111
		}
		goto tr10
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		switch data[p] {
		case 58:
			goto st113
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st112
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st112
			}
		default:
			goto st112
		}
		goto tr10
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 58:
			goto st113
		case 93:
			goto st191
		}
		goto tr10
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		switch data[p] {
		case 48:
			goto st95
		case 49:
			goto st100
		case 50:
			goto st103
		}
		switch {
		case data[p] < 65:
			if 51 <= data[p] && data[p] <= 57 {
				goto st106
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st107
			}
		default:
			goto st107
		}
		goto tr10
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st113
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st115
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st110
			}
		default:
			goto st110
		}
		goto tr10
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st113
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st116
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st111
			}
		default:
			goto st111
		}
		goto tr10
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st113
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st112
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st112
			}
		default:
			goto st112
		}
		goto tr10
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 46:
			goto st46
		case 53:
			goto st118
		case 58:
			goto st113
		case 93:
			goto st191
		}
		switch {
		case data[p] < 54:
			if 48 <= data[p] && data[p] <= 52 {
				goto st115
			}
		case data[p] > 57:
			switch {
			case data[p] > 70:
				if 97 <= data[p] && data[p] <= 102 {
					goto st110
				}
			case data[p] >= 65:
				goto st110
			}
		default:
			goto st119
		}
		goto tr10
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st113
		case 93:
			goto st191
		}
		switch {
		case data[p] < 54:
			if 48 <= data[p] && data[p] <= 53 {
				goto st116
			}
		case data[p] > 57:
			switch {
			case data[p] > 70:
				if 97 <= data[p] && data[p] <= 102 {
					goto st111
				}
			case data[p] >= 65:
				goto st111
			}
		default:
			goto st111
		}
		goto tr10
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st113
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st111
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st111
			}
		default:
			goto st111
		}
		goto tr10
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st113
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st119
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st110
			}
		default:
			goto st110
		}
		goto tr10
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		switch data[p] {
		case 58:
			goto st113
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st110
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st110
			}
		default:
			goto st110
		}
		goto tr10
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 48:
			goto st123
		case 49:
			goto st128
		case 50:
			goto st131
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 51 <= data[p] && data[p] <= 57 {
				goto st134
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st135
			}
		default:
			goto st135
		}
		goto tr10
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st127
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st124
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st124
			}
		default:
			goto st124
		}
		goto tr10
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		switch data[p] {
		case 58:
			goto st127
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st125
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st125
			}
		default:
			goto st125
		}
		goto tr10
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch data[p] {
		case 58:
			goto st127
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st126
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st126
			}
		default:
			goto st126
		}
		goto tr10
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		switch data[p] {
		case 58:
			goto st127
		case 93:
			goto st191
		}
		goto tr10
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		switch data[p] {
		case 48:
			goto st109
		case 49:
			goto st114
		case 50:
			goto st117
		}
		switch {
		case data[p] < 65:
			if 51 <= data[p] && data[p] <= 57 {
				goto st120
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st121
			}
		default:
			goto st121
		}
		goto tr10
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st127
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st129
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st124
			}
		default:
			goto st124
		}
		goto tr10
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st127
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st130
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st125
			}
		default:
			goto st125
		}
		goto tr10
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st127
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st126
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st126
			}
		default:
			goto st126
		}
		goto tr10
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		switch data[p] {
		case 46:
			goto st46
		case 53:
			goto st132
		case 58:
			goto st127
		case 93:
			goto st191
		}
		switch {
		case data[p] < 54:
			if 48 <= data[p] && data[p] <= 52 {
				goto st129
			}
		case data[p] > 57:
			switch {
			case data[p] > 70:
				if 97 <= data[p] && data[p] <= 102 {
					goto st124
				}
			case data[p] >= 65:
				goto st124
			}
		default:
			goto st133
		}
		goto tr10
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st127
		case 93:
			goto st191
		}
		switch {
		case data[p] < 54:
			if 48 <= data[p] && data[p] <= 53 {
				goto st130
			}
		case data[p] > 57:
			switch {
			case data[p] > 70:
				if 97 <= data[p] && data[p] <= 102 {
					goto st125
				}
			case data[p] >= 65:
				goto st125
			}
		default:
			goto st125
		}
		goto tr10
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st127
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st125
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st125
			}
		default:
			goto st125
		}
		goto tr10
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st127
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st133
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st124
			}
		default:
			goto st124
		}
		goto tr10
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		switch data[p] {
		case 58:
			goto st127
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st124
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st124
			}
		default:
			goto st124
		}
		goto tr10
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
		switch data[p] {
		case 48:
			goto st137
		case 49:
			goto st142
		case 50:
			goto st145
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 51 <= data[p] && data[p] <= 57 {
				goto st148
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st149
			}
		default:
			goto st149
		}
		goto tr10
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st141
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st138
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st138
			}
		default:
			goto st138
		}
		goto tr10
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		switch data[p] {
		case 58:
			goto st141
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st139
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st139
			}
		default:
			goto st139
		}
		goto tr10
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		switch data[p] {
		case 58:
			goto st141
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st140
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st140
			}
		default:
			goto st140
		}
		goto tr10
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		switch data[p] {
		case 58:
			goto st141
		case 93:
			goto st191
		}
		goto tr10
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		switch data[p] {
		case 48:
			goto st123
		case 49:
			goto st128
		case 50:
			goto st131
		}
		switch {
		case data[p] < 65:
			if 51 <= data[p] && data[p] <= 57 {
				goto st134
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st135
			}
		default:
			goto st135
		}
		goto tr10
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st141
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st143
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st138
			}
		default:
			goto st138
		}
		goto tr10
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st141
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st144
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st139
			}
		default:
			goto st139
		}
		goto tr10
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st141
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st140
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st140
			}
		default:
			goto st140
		}
		goto tr10
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		switch data[p] {
		case 46:
			goto st46
		case 53:
			goto st146
		case 58:
			goto st141
		case 93:
			goto st191
		}
		switch {
		case data[p] < 54:
			if 48 <= data[p] && data[p] <= 52 {
				goto st143
			}
		case data[p] > 57:
			switch {
			case data[p] > 70:
				if 97 <= data[p] && data[p] <= 102 {
					goto st138
				}
			case data[p] >= 65:
				goto st138
			}
		default:
			goto st147
		}
		goto tr10
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st141
		case 93:
			goto st191
		}
		switch {
		case data[p] < 54:
			if 48 <= data[p] && data[p] <= 53 {
				goto st144
			}
		case data[p] > 57:
			switch {
			case data[p] > 70:
				if 97 <= data[p] && data[p] <= 102 {
					goto st139
				}
			case data[p] >= 65:
				goto st139
			}
		default:
			goto st139
		}
		goto tr10
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st141
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st139
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st139
			}
		default:
			goto st139
		}
		goto tr10
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st141
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st147
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st138
			}
		default:
			goto st138
		}
		goto tr10
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
		switch data[p] {
		case 58:
			goto st141
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st138
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st138
			}
		default:
			goto st138
		}
		goto tr10
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
		if data[p] == 58 {
			goto st151
		}
		goto tr10
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
		switch data[p] {
		case 48:
			goto st152
		case 49:
			goto st157
		case 50:
			goto st160
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 51 <= data[p] && data[p] <= 57 {
				goto st163
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st164
			}
		default:
			goto st164
		}
		goto tr10
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st156
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st153
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st153
			}
		default:
			goto st153
		}
		goto tr10
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		switch data[p] {
		case 58:
			goto st156
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st154
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st154
			}
		default:
			goto st154
		}
		goto tr10
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
		switch data[p] {
		case 58:
			goto st156
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st155
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st155
			}
		default:
			goto st155
		}
		goto tr10
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		switch data[p] {
		case 58:
			goto st156
		case 93:
			goto st191
		}
		goto tr10
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		switch data[p] {
		case 48:
			goto st137
		case 49:
			goto st142
		case 50:
			goto st145
		}
		switch {
		case data[p] < 65:
			if 51 <= data[p] && data[p] <= 57 {
				goto st148
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st149
			}
		default:
			goto st149
		}
		goto tr10
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st156
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st158
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st153
			}
		default:
			goto st153
		}
		goto tr10
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st156
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st159
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st154
			}
		default:
			goto st154
		}
		goto tr10
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st156
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st155
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st155
			}
		default:
			goto st155
		}
		goto tr10
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		switch data[p] {
		case 46:
			goto st46
		case 53:
			goto st161
		case 58:
			goto st156
		case 93:
			goto st191
		}
		switch {
		case data[p] < 54:
			if 48 <= data[p] && data[p] <= 52 {
				goto st158
			}
		case data[p] > 57:
			switch {
			case data[p] > 70:
				if 97 <= data[p] && data[p] <= 102 {
					goto st153
				}
			case data[p] >= 65:
				goto st153
			}
		default:
			goto st162
		}
		goto tr10
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st156
		case 93:
			goto st191
		}
		switch {
		case data[p] < 54:
			if 48 <= data[p] && data[p] <= 53 {
				goto st159
			}
		case data[p] > 57:
			switch {
			case data[p] > 70:
				if 97 <= data[p] && data[p] <= 102 {
					goto st154
				}
			case data[p] >= 65:
				goto st154
			}
		default:
			goto st154
		}
		goto tr10
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st156
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st154
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st154
			}
		default:
			goto st154
		}
		goto tr10
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		switch data[p] {
		case 46:
			goto st46
		case 58:
			goto st156
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st162
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st153
			}
		default:
			goto st153
		}
		goto tr10
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
		switch data[p] {
		case 58:
			goto st156
		case 93:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st153
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st153
			}
		default:
			goto st153
		}
		goto tr10
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st166
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st166
			}
		default:
			goto st166
		}
		goto tr10
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		if data[p] == 46 {
			goto st167
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st166
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st166
			}
		default:
			goto st166
		}
		goto tr10
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		switch data[p] {
		case 33:
			goto st168
		case 36:
			goto st168
		case 61:
			goto st168
		case 95:
			goto st168
		case 126:
			goto st168
		}
		switch {
		case data[p] < 48:
			if 38 <= data[p] && data[p] <= 46 {
				goto st168
			}
		case data[p] > 59:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st168
				}
			case data[p] >= 65:
				goto st168
			}
		default:
			goto st168
		}
		goto tr10
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		switch data[p] {
		case 33:
			goto st168
		case 36:
			goto st168
		case 61:
			goto st168
		case 93:
			goto st191
		case 95:
			goto st168
		case 126:
			goto st168
		}
		switch {
		case data[p] < 48:
			if 38 <= data[p] && data[p] <= 46 {
				goto st168
			}
		case data[p] > 59:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st168
				}
			case data[p] >= 65:
				goto st168
			}
		default:
			goto st168
		}
		goto tr10
	tr7:
//line NONE:1
		te = p + 1

//line yuri.rl:26
		act = 1
		goto st192
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
//line yuri.go:4580
		switch data[p] {
		case 33:
			goto st5
		case 37:
			goto st6
		case 47:
			goto tr19
		case 58:
			goto tr198
		case 60:
			goto tr189
		case 64:
			goto st8
		case 96:
			goto tr189
		case 127:
			goto tr189
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 35:
				if 36 <= data[p] && data[p] <= 44 {
					goto st5
				}
			default:
				goto tr189
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 62 <= data[p] && data[p] <= 63 {
					goto tr189
				}
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 125 {
					goto tr189
				}
			default:
				goto tr189
			}
		default:
			goto st5
		}
		goto st169
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
		switch data[p] {
		case 33:
			goto st5
		case 37:
			goto st6
		case 47:
			goto tr13
		case 60:
			goto tr13
		case 64:
			goto st8
		case 95:
			goto st169
		case 126:
			goto st169
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 36:
				if data[p] <= 35 {
					goto tr13
				}
			case data[p] > 44:
				if 45 <= data[p] && data[p] <= 46 {
					goto st169
				}
			default:
				goto st5
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 62 <= data[p] && data[p] <= 63 {
					goto tr13
				}
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr13
				}
			default:
				goto tr13
			}
		default:
			goto st5
		}
		goto tr176
	tr176:
//line NONE:1
		te = p + 1

//line yuri.rl:26
		act = 1
		goto st193
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
//line yuri.go:4690
		switch data[p] {
		case 33:
			goto st5
		case 37:
			goto st6
		case 47:
			goto tr19
		case 58:
			goto tr198
		case 60:
			goto tr189
		case 64:
			goto st8
		case 95:
			goto st169
		case 126:
			goto st169
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 36:
				if data[p] <= 35 {
					goto tr189
				}
			case data[p] > 44:
				if 45 <= data[p] && data[p] <= 46 {
					goto st169
				}
			default:
				goto st5
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 62 <= data[p] && data[p] <= 63 {
					goto tr189
				}
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr189
				}
			default:
				goto tr189
			}
		default:
			goto st5
		}
		goto tr176
	tr198:
//line NONE:1
		te = p + 1

//line yuri.rl:26
		act = 1
		goto st194
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
//line yuri.go:4752
		switch data[p] {
		case 37:
			goto st6
		case 47:
			goto tr19
		case 60:
			goto tr189
		case 64:
			goto st8
		case 96:
			goto tr189
		case 127:
			goto tr189
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 32:
				if 34 <= data[p] && data[p] <= 35 {
					goto tr189
				}
			default:
				goto tr189
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 62 <= data[p] && data[p] <= 63 {
					goto tr189
				}
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 125 {
					goto tr189
				}
			default:
				goto tr189
			}
		default:
			goto tr199
		}
		goto st5
	tr199:
//line NONE:1
		te = p + 1

//line yuri.rl:26
		act = 1
		goto st195
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
//line yuri.go:4806
		switch data[p] {
		case 37:
			goto st6
		case 47:
			goto tr19
		case 60:
			goto tr189
		case 64:
			goto st8
		case 96:
			goto tr189
		case 127:
			goto tr189
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 32:
				if 34 <= data[p] && data[p] <= 35 {
					goto tr189
				}
			default:
				goto tr189
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 62 <= data[p] && data[p] <= 63 {
					goto tr189
				}
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 125 {
					goto tr189
				}
			default:
				goto tr189
			}
		default:
			goto tr200
		}
		goto st5
	tr200:
//line NONE:1
		te = p + 1

//line yuri.rl:26
		act = 1
		goto st196
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
//line yuri.go:4860
		switch data[p] {
		case 37:
			goto st6
		case 47:
			goto tr19
		case 60:
			goto tr189
		case 64:
			goto st8
		case 96:
			goto tr189
		case 127:
			goto tr189
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 32:
				if 34 <= data[p] && data[p] <= 35 {
					goto tr189
				}
			default:
				goto tr189
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 62 <= data[p] && data[p] <= 63 {
					goto tr189
				}
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 125 {
					goto tr189
				}
			default:
				goto tr189
			}
		default:
			goto tr201
		}
		goto st5
	tr201:
//line NONE:1
		te = p + 1

//line yuri.rl:26
		act = 1
		goto st197
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
//line yuri.go:4914
		switch data[p] {
		case 37:
			goto st6
		case 47:
			goto tr19
		case 60:
			goto tr189
		case 64:
			goto st8
		case 96:
			goto tr189
		case 127:
			goto tr189
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 32:
				if 34 <= data[p] && data[p] <= 35 {
					goto tr189
				}
			default:
				goto tr189
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 62 <= data[p] && data[p] <= 63 {
					goto tr189
				}
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 125 {
					goto tr189
				}
			default:
				goto tr189
			}
		default:
			goto tr202
		}
		goto st5
	tr202:
//line NONE:1
		te = p + 1

//line yuri.rl:26
		act = 1
		goto st198
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
//line yuri.go:4968
		switch data[p] {
		case 37:
			goto st6
		case 47:
			goto tr19
		case 60:
			goto tr189
		case 64:
			goto st8
		case 96:
			goto tr189
		case 127:
			goto tr189
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 32:
				if 34 <= data[p] && data[p] <= 35 {
					goto tr189
				}
			default:
				goto tr189
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 62 <= data[p] && data[p] <= 63 {
					goto tr189
				}
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 125 {
					goto tr189
				}
			default:
				goto tr189
			}
		default:
			goto tr203
		}
		goto st5
	tr203:
//line NONE:1
		te = p + 1

//line yuri.rl:26
		act = 1
		goto st199
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
//line yuri.go:5022
		switch data[p] {
		case 37:
			goto st6
		case 47:
			goto tr19
		case 60:
			goto tr189
		case 64:
			goto st8
		case 96:
			goto tr189
		case 127:
			goto tr189
		}
		switch {
		case data[p] < 62:
			switch {
			case data[p] > 32:
				if 34 <= data[p] && data[p] <= 35 {
					goto tr189
				}
			default:
				goto tr189
			}
		case data[p] > 63:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 125 {
					goto tr189
				}
			case data[p] >= 91:
				goto tr189
			}
		default:
			goto tr189
		}
		goto st5
	tr185:
//line NONE:1
		te = p + 1

//line yuri.rl:82
		act = 2
		goto st200
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
//line yuri.go:5072
		switch data[p] {
		case 84:
			goto st170
		case 88:
			goto st173
		case 116:
			goto st170
		case 120:
			goto st173
		}
		goto tr187
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		switch data[p] {
		case 84:
			goto st171
		case 116:
			goto st171
		}
		goto tr0
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		switch data[p] {
		case 80:
			goto st172
		case 112:
			goto st172
		}
		goto tr0
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		switch data[p] {
		case 58:
			goto st2
		case 83:
			goto st1
		case 115:
			goto st1
		}
		goto tr0
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
		switch data[p] {
		case 88:
			goto st171
		case 120:
			goto st171
		}
		goto tr0
	tr186:
//line NONE:1
		te = p + 1

//line yuri.rl:82
		act = 2
		goto st201
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
//line yuri.go:5146
		switch data[p] {
		case 65:
			goto st174
		case 97:
			goto st174
		}
		goto tr187
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		switch data[p] {
		case 73:
			goto st175
		case 105:
			goto st175
		}
		goto tr0
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
		switch data[p] {
		case 76:
			goto st176
		case 108:
			goto st176
		}
		goto tr0
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
		switch data[p] {
		case 84:
			goto st177
		case 116:
			goto st177
		}
		goto tr0
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
		switch data[p] {
		case 79:
			goto st178
		case 111:
			goto st178
		}
		goto tr0
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		if data[p] == 58 {
			goto st4
		}
		goto tr0
	st_out:
	_test_eof179:
		cs = 179
		goto _test_eof
	_test_eof180:
		cs = 180
		goto _test_eof
	_test_eof0:
		cs = 0
		goto _test_eof
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof181:
		cs = 181
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof182:
		cs = 182
		goto _test_eof
	_test_eof183:
		cs = 183
		goto _test_eof
	_test_eof184:
		cs = 184
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof185:
		cs = 185
		goto _test_eof
	_test_eof186:
		cs = 186
		goto _test_eof
	_test_eof187:
		cs = 187
		goto _test_eof
	_test_eof188:
		cs = 188
		goto _test_eof
	_test_eof189:
		cs = 189
		goto _test_eof
	_test_eof190:
		cs = 190
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof
	_test_eof31:
		cs = 31
		goto _test_eof
	_test_eof32:
		cs = 32
		goto _test_eof
	_test_eof33:
		cs = 33
		goto _test_eof
	_test_eof34:
		cs = 34
		goto _test_eof
	_test_eof35:
		cs = 35
		goto _test_eof
	_test_eof36:
		cs = 36
		goto _test_eof
	_test_eof37:
		cs = 37
		goto _test_eof
	_test_eof38:
		cs = 38
		goto _test_eof
	_test_eof39:
		cs = 39
		goto _test_eof
	_test_eof40:
		cs = 40
		goto _test_eof
	_test_eof41:
		cs = 41
		goto _test_eof
	_test_eof42:
		cs = 42
		goto _test_eof
	_test_eof43:
		cs = 43
		goto _test_eof
	_test_eof44:
		cs = 44
		goto _test_eof
	_test_eof45:
		cs = 45
		goto _test_eof
	_test_eof46:
		cs = 46
		goto _test_eof
	_test_eof47:
		cs = 47
		goto _test_eof
	_test_eof48:
		cs = 48
		goto _test_eof
	_test_eof49:
		cs = 49
		goto _test_eof
	_test_eof50:
		cs = 50
		goto _test_eof
	_test_eof51:
		cs = 51
		goto _test_eof
	_test_eof191:
		cs = 191
		goto _test_eof
	_test_eof52:
		cs = 52
		goto _test_eof
	_test_eof53:
		cs = 53
		goto _test_eof
	_test_eof54:
		cs = 54
		goto _test_eof
	_test_eof55:
		cs = 55
		goto _test_eof
	_test_eof56:
		cs = 56
		goto _test_eof
	_test_eof57:
		cs = 57
		goto _test_eof
	_test_eof58:
		cs = 58
		goto _test_eof
	_test_eof59:
		cs = 59
		goto _test_eof
	_test_eof60:
		cs = 60
		goto _test_eof
	_test_eof61:
		cs = 61
		goto _test_eof
	_test_eof62:
		cs = 62
		goto _test_eof
	_test_eof63:
		cs = 63
		goto _test_eof
	_test_eof64:
		cs = 64
		goto _test_eof
	_test_eof65:
		cs = 65
		goto _test_eof
	_test_eof66:
		cs = 66
		goto _test_eof
	_test_eof67:
		cs = 67
		goto _test_eof
	_test_eof68:
		cs = 68
		goto _test_eof
	_test_eof69:
		cs = 69
		goto _test_eof
	_test_eof70:
		cs = 70
		goto _test_eof
	_test_eof71:
		cs = 71
		goto _test_eof
	_test_eof72:
		cs = 72
		goto _test_eof
	_test_eof73:
		cs = 73
		goto _test_eof
	_test_eof74:
		cs = 74
		goto _test_eof
	_test_eof75:
		cs = 75
		goto _test_eof
	_test_eof76:
		cs = 76
		goto _test_eof
	_test_eof77:
		cs = 77
		goto _test_eof
	_test_eof78:
		cs = 78
		goto _test_eof
	_test_eof79:
		cs = 79
		goto _test_eof
	_test_eof80:
		cs = 80
		goto _test_eof
	_test_eof81:
		cs = 81
		goto _test_eof
	_test_eof82:
		cs = 82
		goto _test_eof
	_test_eof83:
		cs = 83
		goto _test_eof
	_test_eof84:
		cs = 84
		goto _test_eof
	_test_eof85:
		cs = 85
		goto _test_eof
	_test_eof86:
		cs = 86
		goto _test_eof
	_test_eof87:
		cs = 87
		goto _test_eof
	_test_eof88:
		cs = 88
		goto _test_eof
	_test_eof89:
		cs = 89
		goto _test_eof
	_test_eof90:
		cs = 90
		goto _test_eof
	_test_eof91:
		cs = 91
		goto _test_eof
	_test_eof92:
		cs = 92
		goto _test_eof
	_test_eof93:
		cs = 93
		goto _test_eof
	_test_eof94:
		cs = 94
		goto _test_eof
	_test_eof95:
		cs = 95
		goto _test_eof
	_test_eof96:
		cs = 96
		goto _test_eof
	_test_eof97:
		cs = 97
		goto _test_eof
	_test_eof98:
		cs = 98
		goto _test_eof
	_test_eof99:
		cs = 99
		goto _test_eof
	_test_eof100:
		cs = 100
		goto _test_eof
	_test_eof101:
		cs = 101
		goto _test_eof
	_test_eof102:
		cs = 102
		goto _test_eof
	_test_eof103:
		cs = 103
		goto _test_eof
	_test_eof104:
		cs = 104
		goto _test_eof
	_test_eof105:
		cs = 105
		goto _test_eof
	_test_eof106:
		cs = 106
		goto _test_eof
	_test_eof107:
		cs = 107
		goto _test_eof
	_test_eof108:
		cs = 108
		goto _test_eof
	_test_eof109:
		cs = 109
		goto _test_eof
	_test_eof110:
		cs = 110
		goto _test_eof
	_test_eof111:
		cs = 111
		goto _test_eof
	_test_eof112:
		cs = 112
		goto _test_eof
	_test_eof113:
		cs = 113
		goto _test_eof
	_test_eof114:
		cs = 114
		goto _test_eof
	_test_eof115:
		cs = 115
		goto _test_eof
	_test_eof116:
		cs = 116
		goto _test_eof
	_test_eof117:
		cs = 117
		goto _test_eof
	_test_eof118:
		cs = 118
		goto _test_eof
	_test_eof119:
		cs = 119
		goto _test_eof
	_test_eof120:
		cs = 120
		goto _test_eof
	_test_eof121:
		cs = 121
		goto _test_eof
	_test_eof122:
		cs = 122
		goto _test_eof
	_test_eof123:
		cs = 123
		goto _test_eof
	_test_eof124:
		cs = 124
		goto _test_eof
	_test_eof125:
		cs = 125
		goto _test_eof
	_test_eof126:
		cs = 126
		goto _test_eof
	_test_eof127:
		cs = 127
		goto _test_eof
	_test_eof128:
		cs = 128
		goto _test_eof
	_test_eof129:
		cs = 129
		goto _test_eof
	_test_eof130:
		cs = 130
		goto _test_eof
	_test_eof131:
		cs = 131
		goto _test_eof
	_test_eof132:
		cs = 132
		goto _test_eof
	_test_eof133:
		cs = 133
		goto _test_eof
	_test_eof134:
		cs = 134
		goto _test_eof
	_test_eof135:
		cs = 135
		goto _test_eof
	_test_eof136:
		cs = 136
		goto _test_eof
	_test_eof137:
		cs = 137
		goto _test_eof
	_test_eof138:
		cs = 138
		goto _test_eof
	_test_eof139:
		cs = 139
		goto _test_eof
	_test_eof140:
		cs = 140
		goto _test_eof
	_test_eof141:
		cs = 141
		goto _test_eof
	_test_eof142:
		cs = 142
		goto _test_eof
	_test_eof143:
		cs = 143
		goto _test_eof
	_test_eof144:
		cs = 144
		goto _test_eof
	_test_eof145:
		cs = 145
		goto _test_eof
	_test_eof146:
		cs = 146
		goto _test_eof
	_test_eof147:
		cs = 147
		goto _test_eof
	_test_eof148:
		cs = 148
		goto _test_eof
	_test_eof149:
		cs = 149
		goto _test_eof
	_test_eof150:
		cs = 150
		goto _test_eof
	_test_eof151:
		cs = 151
		goto _test_eof
	_test_eof152:
		cs = 152
		goto _test_eof
	_test_eof153:
		cs = 153
		goto _test_eof
	_test_eof154:
		cs = 154
		goto _test_eof
	_test_eof155:
		cs = 155
		goto _test_eof
	_test_eof156:
		cs = 156
		goto _test_eof
	_test_eof157:
		cs = 157
		goto _test_eof
	_test_eof158:
		cs = 158
		goto _test_eof
	_test_eof159:
		cs = 159
		goto _test_eof
	_test_eof160:
		cs = 160
		goto _test_eof
	_test_eof161:
		cs = 161
		goto _test_eof
	_test_eof162:
		cs = 162
		goto _test_eof
	_test_eof163:
		cs = 163
		goto _test_eof
	_test_eof164:
		cs = 164
		goto _test_eof
	_test_eof165:
		cs = 165
		goto _test_eof
	_test_eof166:
		cs = 166
		goto _test_eof
	_test_eof167:
		cs = 167
		goto _test_eof
	_test_eof168:
		cs = 168
		goto _test_eof
	_test_eof192:
		cs = 192
		goto _test_eof
	_test_eof169:
		cs = 169
		goto _test_eof
	_test_eof193:
		cs = 193
		goto _test_eof
	_test_eof194:
		cs = 194
		goto _test_eof
	_test_eof195:
		cs = 195
		goto _test_eof
	_test_eof196:
		cs = 196
		goto _test_eof
	_test_eof197:
		cs = 197
		goto _test_eof
	_test_eof198:
		cs = 198
		goto _test_eof
	_test_eof199:
		cs = 199
		goto _test_eof
	_test_eof200:
		cs = 200
		goto _test_eof
	_test_eof170:
		cs = 170
		goto _test_eof
	_test_eof171:
		cs = 171
		goto _test_eof
	_test_eof172:
		cs = 172
		goto _test_eof
	_test_eof173:
		cs = 173
		goto _test_eof
	_test_eof201:
		cs = 201
		goto _test_eof
	_test_eof174:
		cs = 174
		goto _test_eof
	_test_eof175:
		cs = 175
		goto _test_eof
	_test_eof176:
		cs = 176
		goto _test_eof
	_test_eof177:
		cs = 177
		goto _test_eof
	_test_eof178:
		cs = 178
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 180:
				goto tr187
			case 0:
				goto tr0
			case 1:
				goto tr0
			case 2:
				goto tr0
			case 3:
				goto tr0
			case 4:
				goto tr0
			case 5:
				goto tr10
			case 6:
				goto tr10
			case 7:
				goto tr10
			case 8:
				goto tr10
			case 181:
				goto tr189
			case 9:
				goto tr13
			case 182:
				goto tr189
			case 183:
				goto tr189
			case 184:
				goto tr189
			case 10:
				goto tr13
			case 11:
				goto tr13
			case 12:
				goto tr13
			case 13:
				goto tr13
			case 185:
				goto tr189
			case 186:
				goto tr189
			case 187:
				goto tr189
			case 188:
				goto tr189
			case 189:
				goto tr189
			case 190:
				goto tr189
			case 14:
				goto tr10
			case 15:
				goto tr10
			case 16:
				goto tr10
			case 17:
				goto tr10
			case 18:
				goto tr10
			case 19:
				goto tr10
			case 20:
				goto tr10
			case 21:
				goto tr10
			case 22:
				goto tr10
			case 23:
				goto tr10
			case 24:
				goto tr10
			case 25:
				goto tr10
			case 26:
				goto tr10
			case 27:
				goto tr10
			case 28:
				goto tr10
			case 29:
				goto tr10
			case 30:
				goto tr10
			case 31:
				goto tr10
			case 32:
				goto tr10
			case 33:
				goto tr10
			case 34:
				goto tr10
			case 35:
				goto tr10
			case 36:
				goto tr10
			case 37:
				goto tr10
			case 38:
				goto tr10
			case 39:
				goto tr10
			case 40:
				goto tr10
			case 41:
				goto tr10
			case 42:
				goto tr10
			case 43:
				goto tr10
			case 44:
				goto tr10
			case 45:
				goto tr10
			case 46:
				goto tr10
			case 47:
				goto tr10
			case 48:
				goto tr10
			case 49:
				goto tr10
			case 50:
				goto tr10
			case 51:
				goto tr10
			case 191:
				goto tr189
			case 52:
				goto tr10
			case 53:
				goto tr10
			case 54:
				goto tr10
			case 55:
				goto tr10
			case 56:
				goto tr10
			case 57:
				goto tr10
			case 58:
				goto tr10
			case 59:
				goto tr10
			case 60:
				goto tr10
			case 61:
				goto tr10
			case 62:
				goto tr10
			case 63:
				goto tr10
			case 64:
				goto tr10
			case 65:
				goto tr10
			case 66:
				goto tr10
			case 67:
				goto tr10
			case 68:
				goto tr10
			case 69:
				goto tr10
			case 70:
				goto tr10
			case 71:
				goto tr10
			case 72:
				goto tr10
			case 73:
				goto tr10
			case 74:
				goto tr10
			case 75:
				goto tr10
			case 76:
				goto tr10
			case 77:
				goto tr10
			case 78:
				goto tr10
			case 79:
				goto tr10
			case 80:
				goto tr10
			case 81:
				goto tr10
			case 82:
				goto tr10
			case 83:
				goto tr10
			case 84:
				goto tr10
			case 85:
				goto tr10
			case 86:
				goto tr10
			case 87:
				goto tr10
			case 88:
				goto tr10
			case 89:
				goto tr10
			case 90:
				goto tr10
			case 91:
				goto tr10
			case 92:
				goto tr10
			case 93:
				goto tr10
			case 94:
				goto tr10
			case 95:
				goto tr10
			case 96:
				goto tr10
			case 97:
				goto tr10
			case 98:
				goto tr10
			case 99:
				goto tr10
			case 100:
				goto tr10
			case 101:
				goto tr10
			case 102:
				goto tr10
			case 103:
				goto tr10
			case 104:
				goto tr10
			case 105:
				goto tr10
			case 106:
				goto tr10
			case 107:
				goto tr10
			case 108:
				goto tr10
			case 109:
				goto tr10
			case 110:
				goto tr10
			case 111:
				goto tr10
			case 112:
				goto tr10
			case 113:
				goto tr10
			case 114:
				goto tr10
			case 115:
				goto tr10
			case 116:
				goto tr10
			case 117:
				goto tr10
			case 118:
				goto tr10
			case 119:
				goto tr10
			case 120:
				goto tr10
			case 121:
				goto tr10
			case 122:
				goto tr10
			case 123:
				goto tr10
			case 124:
				goto tr10
			case 125:
				goto tr10
			case 126:
				goto tr10
			case 127:
				goto tr10
			case 128:
				goto tr10
			case 129:
				goto tr10
			case 130:
				goto tr10
			case 131:
				goto tr10
			case 132:
				goto tr10
			case 133:
				goto tr10
			case 134:
				goto tr10
			case 135:
				goto tr10
			case 136:
				goto tr10
			case 137:
				goto tr10
			case 138:
				goto tr10
			case 139:
				goto tr10
			case 140:
				goto tr10
			case 141:
				goto tr10
			case 142:
				goto tr10
			case 143:
				goto tr10
			case 144:
				goto tr10
			case 145:
				goto tr10
			case 146:
				goto tr10
			case 147:
				goto tr10
			case 148:
				goto tr10
			case 149:
				goto tr10
			case 150:
				goto tr10
			case 151:
				goto tr10
			case 152:
				goto tr10
			case 153:
				goto tr10
			case 154:
				goto tr10
			case 155:
				goto tr10
			case 156:
				goto tr10
			case 157:
				goto tr10
			case 158:
				goto tr10
			case 159:
				goto tr10
			case 160:
				goto tr10
			case 161:
				goto tr10
			case 162:
				goto tr10
			case 163:
				goto tr10
			case 164:
				goto tr10
			case 165:
				goto tr10
			case 166:
				goto tr10
			case 167:
				goto tr10
			case 168:
				goto tr10
			case 192:
				goto tr189
			case 169:
				goto tr13
			case 193:
				goto tr189
			case 194:
				goto tr189
			case 195:
				goto tr189
			case 196:
				goto tr189
			case 197:
				goto tr189
			case 198:
				goto tr189
			case 199:
				goto tr189
			case 200:
				goto tr187
			case 170:
				goto tr0
			case 171:
				goto tr0
			case 172:
				goto tr0
			case 173:
				goto tr0
			case 201:
				goto tr187
			case 174:
				goto tr0
			case 175:
				goto tr0
			case 176:
				goto tr0
			case 177:
				goto tr0
			case 178:
				goto tr0
			}
		}

	}

//line yuri.rl:88

	return uris
}
