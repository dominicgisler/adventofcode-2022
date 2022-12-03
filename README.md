# Advent of Code 2022

[https://adventofcode.com/2022](https://adventofcode.com/2022)

<style>
.calendar {
  background: #0f0f23;
}
.calendar > span {
  color: #333333;
}
.calendar > a {
  text-decoration: none;
  color: #666666;
  outline: none;
  cursor: default;
}
.calendar a:hover, .calendar a:focus {
  background-color: #1e1e46;
  background-color: rgba(119,119,165,.2);
  cursor: pointer;
}
.calendar .calendar-day { color: #666666; }
.calendar a .calendar-day { color: #cccccc; }
.calendar a .calendar-mark-complete,
.calendar a .calendar-mark-verycomplete { visibility: hidden; }
.calendar a.calendar-complete     .calendar-mark-complete,
.calendar a.calendar-verycomplete .calendar-mark-complete { visibility: visible; color: #ffff66; }
.calendar a.calendar-verycomplete .calendar-mark-verycomplete { visibility: visible; color: #ffff66; }
.calendar .calendar-color-g2 { color:#7fbd39; }
.calendar .calendar-color-u { color:#5eabb4; }
.calendar .calendar-color-g3 { color:#427322; }
.calendar .calendar-color-s { color:#d0b376; }
.calendar .calendar-color-g1 { color:#4d8b03; }
</style>

<pre class="calendar"><span aria-hidden="true" class="calendar-day25">@@@@@@@@@@@@@@@#@@@@@@#@@@@@#@#@@@@@@##@@@@@@##@#  <span class="calendar-day">25</span></span>
<span aria-hidden="true" class="calendar-day24">@###@@@#@@#@@@#@@@@@@@@@@#@@@@@@@##@#@@@@#@@#@@@@  <span class="calendar-day">24</span></span>
<span aria-hidden="true" class="calendar-day23">@@#@@#@@@@#@@##@@@@@@#@##@@#|@@@#@#@##@@@#@##@@#@  <span class="calendar-day">23</span></span>
<span aria-hidden="true" class="calendar-day22">####@@@#@@#@#@@@@@@##@@@@#@#@#@@@@@@@@###@@@###|@  <span class="calendar-day">22</span></span>
<span aria-hidden="true" class="calendar-day21">@@@@@@#@@@@@@#@@#@#@@@@@@#@@#@#@@@#@@#@@#@#@@#@@@  <span class="calendar-day">21</span></span>
<span aria-hidden="true" class="calendar-day20">@@@@@@@@@@#@###@@@@@#@@@@@@@@#@@#@#@@@@@@@@@##@@@  <span class="calendar-day">20</span></span>
<span aria-hidden="true" class="calendar-day19">#@@##@@@#@@@#@#@@@@@@@@@@@@#@@@@@@#@@@@@@@@@@@@@@  <span class="calendar-day">19</span></span>
<span aria-hidden="true" class="calendar-day18">@@@#@@@#@#@@#@@@@@@@##@@@##@@@#@##@#@#@@@@#@@#@@#  <span class="calendar-day">18</span></span>
<span aria-hidden="true" class="calendar-day17">@@#@@#@@#@@@#@@@@@@####@###@@@#@@@@@#|@@@@@#@##@@  <span class="calendar-day">17</span></span>
<span aria-hidden="true" class="calendar-day16">@#@@@@@##@@@#@@@@#@@#@@@#@@#@@#@@@@@#@@@##@@#@@@#  <span class="calendar-day">16</span></span>
<span aria-hidden="true" class="calendar-day15">##@#@@@@@@####@@###@###@@#@@@#@@@@@@@###@@#@##@@@  <span class="calendar-day">15</span></span>
<span aria-hidden="true" class="calendar-day14">##@##@@#@@@##@@@@#@@@@#@@@#@###@##@@@@@#@@@@#@@@@  <span class="calendar-day">14</span></span>
<span aria-hidden="true" class="calendar-day13">@@@@@@@@@##@@###@@#@@@#@@##@@#@@@@@@@#@@@@@#@##@#  <span class="calendar-day">13</span></span>
<span aria-hidden="true" class="calendar-day12">#@@@#@#@@@##@@@@@#@@@@###@##@#@@#@@@@#@##@@@#@@#@  <span class="calendar-day">12</span></span>
<span aria-hidden="true" class="calendar-day11">@#@@@@##@@@@@@#@#@@##@#@@@@@@##@@@@@@#@@##@#@###@  <span class="calendar-day">11</span></span>
<span aria-hidden="true" class="calendar-day10">##@#@@@@@@@@@##@#@@@###@#@@##@@@#@@@@#@@@@@@#@##@  <span class="calendar-day">10</span></span>
<span aria-hidden="true" class="calendar-day9">@@@@@@@#@@#@#@@@#@@#@#@#@@@##@@@@#@@@##@#@######@  <span class="calendar-day"> 9</span></span>
<span aria-hidden="true" class="calendar-day8">@@#@@@@@#@#@@@@#@@@@@#@#@#@@@@@@@@@@@@##@@@@#@@@@  <span class="calendar-day"> 8</span></span>
<span aria-hidden="true" class="calendar-day7">@@@@#@#@#@@@##@@@@#@##@@@##@#@@@@#@##@##@@#@###@@  <span class="calendar-day"> 7</span></span>
<span aria-hidden="true" class="calendar-day6">@###@##@@@@#@@#@@@@#@@@@@@@#@#@@@#@@@@@@@@###@@@#  <span class="calendar-day"> 6</span></span>
<span aria-hidden="true" class="calendar-day5">@@@@@#@@@@@@@@@@@@#@#@@@@@@@##@@#@@#@####@@@@@@#@  <span class="calendar-day"> 5</span></span>
<span aria-hidden="true" class="calendar-day4">@@@#@@@@@@@@@@@@@@@@@@@@@@#@@@@@@@@@@#@#@@@@@@@@@  <span class="calendar-day"> 4</span></span>
<a aria-label="Day 3, two stars" href="day3" class="calendar-day3 calendar-verycomplete">@@@@@@#<span class="calendar-color-g2">@</span><span class="calendar-color-g3">@</span><span class="calendar-color-s">_/</span><span class="calendar-color-u"> ~   ~  </span><span class="calendar-color-s">\ ' '. '.'.</span><span class="calendar-color-g3">#</span><span class="calendar-color-g1">@</span>@##@@#@##@@@@@@@@  <span class="calendar-day"> 3</span> <span class="calendar-mark-complete">*</span><span class="calendar-mark-verycomplete">*</span></a>
<a aria-label="Day 2, two stars" href="day2" class="calendar-day2 calendar-verycomplete"><span class="calendar-color-s">-~------'</span><span class="calendar-color-u">    ~    ~ </span><span class="calendar-color-s">'--~-----~-~----___________--</span>  <span class="calendar-day"> 2</span> <span class="calendar-mark-complete">*</span><span class="calendar-mark-verycomplete">*</span></a>
<a aria-label="Day 1, two stars" href="day1" class="calendar-day1 calendar-verycomplete"><span class="calendar-color-u">  ~    ~  ~      ~     ~ ~   ~     ~  ~  ~   ~   </span>  <span class="calendar-day"> 1</span> <span class="calendar-mark-complete">*</span><span class="calendar-mark-verycomplete">*</span></a>
</pre>
