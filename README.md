# calm

A tiny terminal breathing companion. A single soft circle expands and contracts
in your terminal — inhale as it grows, exhale as it shrinks. Stare and breathe.

![calm in coherent mode: a circle of dots slowly grows and shrinks](docs/demo.gif)

No config, no dependencies, no network. Just one Go file.

## Install

```sh
go install github.com/jmserra/calm@latest
```

Or from a checkout:

```sh
make install        # builds and copies to ~/.local/bin (override with BINDIR=...)
```

Or build it yourself:

```sh
make build          # produces ./calm
go build .          # equivalent
```

Requires Go 1.24+. Works on any terminal that understands 256-color ANSI
escapes, on Linux or macOS. It does not build on Windows.

## Usage

```sh
calm [mode]         # default: coherent
calm --help         # list modes and their timing
```

The circle animates at ~30fps. Press **ctrl-c** to leave — the terminal is
restored cleanly on exit.

## Modes

Timing is `inhale-hold-exhale-hold`, in seconds.

| mode       | timing        | what it's for |
|------------|---------------|---------------|
| `coherent` | 5.5-0-5.5-0   | ~6 breaths/min. Resonant, evidence-backed default for just relaxing. *(default)* |
| `box`      | 4-4-4-4       | Equal four counts. Grounding under stress; used by military/first responders. |
| `478`      | 4-7-8-0       | Dr. Weil's sedative pattern. Good for falling asleep or coming down from anxiety. |
| `exhale`   | 4-0-6-0       | Longer out than in. A gentle parasympathetic nudge, no holds to think about. |

## How it works

Each breathing pattern is four durations. A cosine ease maps the current time
within the cycle to a radius fraction (0 collapsed, 1 full), holds clamp to the
endpoints, and `draw` fills every terminal cell within that radius — squashing
the vertical axis so the circle reads round despite tall character cells. The
radius never collapses fully, so there's always something to rest your eyes on.

## The science

Slow, controlled breathing is one of the few voluntary levers we have on the
autonomic nervous system. The evidence base is real but uneven: strongest for
slow-paced/coherent breathing and exhale-emphasized practices, thinner and more
marketing-driven for named protocols like box breathing and 4-7-8. Here is what
the literature actually shows — and where it doesn't.

### Coherent / resonance breathing (~6 breaths/min) — the `coherent` mode

Breathing at roughly 5.5–6 breaths per minute (about 0.1 Hz) hits the
**resonance frequency** of the cardiovascular system. At this rate the rhythms
of respiration, heart rate, and blood pressure line up in phase, producing
unusually large heart-rate oscillations — high **heart-rate variability (HRV)**
— and strongly engaging the **baroreflex**, the loop that regulates blood
pressure. The mechanism and clinical groundwork come from Evgeny Vaschillo and
Paul Lehrer, pioneers of HRV biofeedback (Lehrer et al., 2022). A large review,
*"HRV and slow-paced breathing: when coherence meets resonance"* (Sevoz-Couche &
Laborde, 2022), concludes slow-paced breathing shows substantial efficacy for
reducing subjective stress and anxiety via vagal/baroreflex pathways.
Individual resonance frequency typically lands between 4.5 and 6.5 breaths/min
(Shaffer & Meehan, 2020). One honest caveat: a placebo-controlled trial
(Fincham et al., 2023) found coherent breathing did *not* significantly
outperform a matched placebo on mental wellbeing — controlled designs sometimes
shrink the effect.

### Box breathing (4-4-4-4) — the `box` mode

Equal-count inhale, hold, exhale, hold. Popularized as "tactical breathing" in
U.S. military and law-enforcement training, and promoted by former Navy SEAL
Mark Divine, who frames the 4-4-4-4 ratio as energetically neutral. Direct
evidence for box breathing *specifically* is thin — its best controlled
appearance is as one arm of the Stanford trial below, where it helped but was
outperformed by exhale-focused breathing. At ~3.75 breaths/min it still sits
well within the slow-breathing range and plausibly works through the same
vagal/baroreflex mechanisms. Honestly: the military pedigree is a story, not
evidence.

### 4-7-8 breathing — the `478` mode

Inhale 4, hold 7, exhale 8 — a long hold and a markedly extended exhale.
Popularized by Dr. Andrew Weil from yogic pranayama. Studied directly, but
mostly in small trials in specific populations (nursing students, COPD
patients) and one study of 43 healthy young adults (Vierra et al., 2022) showing
acute drops in heart rate and blood pressure and a shift toward parasympathetic
dominance in HRV. Several studies report reduced state anxiety versus controls.
There is **no large, high-quality RCT** behind the popular "fall asleep in 60
seconds" claim. The 8-count exhale is the active ingredient.

### Extended-exhale breathing — the `exhale` mode

Making the exhale longer than the inhale is the most physiologically
well-justified pattern for relaxation. The cardiac cycle is asymmetric: during
inhalation sympathetic influence briefly speeds the heart; during exhalation the
vagus nerve releases acetylcholine and the heart slows. Lengthening the exhale
biases autonomic balance toward **parasympathetic dominance**. Komori (2018)
found prolonged expiration significantly activated parasympathetic function;
Magnon et al. (2021) found 5 minutes of slow breathing with a longer exhale
increased parasympathetic activity and reduced anxiety in young and older
adults. Nuance: during *already-slow* paced breathing, lengthening the exhale
further may not raise HRV more (2024, *Appl. Psychophysiol. Biofeedback*) — the
long-exhale benefit is clearest relative to normal-rate breathing.

### Slow breathing generally — what underpins all four modes

Across techniques, voluntarily slowing the breath below ~10 breaths/min produces
a consistent cluster of effects: a meta-analysis (Zhang et al., 2023) found
significant immediate reductions in systolic blood pressure and improvements in
emotional measures; controlled studies of diaphragmatic breathing report lower
post-intervention salivary cortisol (Ma et al., 2017); sustained programs lower
resting heart rate and blood pressure. The mechanism is increased vagal tone and
baroreflex sensitivity, dampened sympathetic outflow, and apparently lower
HPA-axis (cortisol) activity.

### Two meta-analyses worth knowing

- **Fincham, Strauss, Montero-Marin & Cavanagh (2023), *Scientific Reports*** —
  meta-analysis of 12 RCTs (785 adults). Breathwork showed a small-to-medium
  effect on stress (Hedges *g* ≈ −0.35), with similar significant effects for
  anxiety and depression. The authors explicitly urge caution "to avoid a
  miscalibration between hype and evidence."
- **Balban, Huberman, Spiegel et al. (2023), *Cell Reports Medicine*** — remote
  RCT, 111 participants, comparing three 5-minute/day breathwork practices
  against mindfulness meditation for a month. All helped, but breathwork — and
  especially **cyclic sighing** (long exhales) — improved positive affect and
  reduced respiratory rate more than meditation. Exhale-weighted practices have a
  slight edge.

### Caveats — read this before believing the hype

- **Evidence quality is modest.** Most positive findings come from small trials,
  often in narrow populations, with moderate risk of bias and short follow-up.
  Effect sizes are small-to-medium, not dramatic.
- **Named protocols are under-studied.** "Box breathing" and "4-7-8" are largely
  brand names; the rigorous evidence is for slow-paced and exhale-emphasized
  breathing in general, not the specific counts.
- **Breathwork is an adjunct, not a treatment.** It is not a substitute for care
  for diagnosed anxiety, depression, hypertension, or any medical condition.
- **Contraindications.** The gentle practices `calm` offers — slow breathing at
  ~6 breaths/min with mild or no holds — have a favorable safety profile for
  healthy adults. But seek medical advice before doing breath holds or any
  forceful breathing if you have cardiovascular disease, asthma or other
  respiratory conditions, a seizure history, panic disorder or PTSD, or are
  pregnant. Never do breath holds in or before water. If you feel dizzy or
  tingly, stop and breathe normally.

### References

1. Fincham GW, Strauss C, Montero-Marin J, Cavanagh K. *Effect of breathwork on stress and mental health: A meta-analysis of randomised-controlled trials.* Scientific Reports, 2023. <https://www.nature.com/articles/s41598-022-27247-y>
2. Balban MY, Neri E, Kogon MM, et al. *Brief structured respiration practices enhance mood and reduce physiological arousal.* Cell Reports Medicine, 2023. <https://www.cell.com/cell-reports-medicine/fulltext/S2666-3791(22)00474-8>
3. Sevoz-Couche C, Laborde S. *Heart rate variability and slow-paced breathing: when coherence meets resonance.* Neuroscience & Biobehavioral Reviews, 2022. <https://pubmed.ncbi.nlm.nih.gov/35167847/>
4. Lehrer P, et al. *The Process of HRV, Resonance at 0.1 Hz, and the Three Baroreflex Loops: A Tribute to Evgeny Vaschillo.* Applied Psychophysiology and Biofeedback, 2022. <https://pmc.ncbi.nlm.nih.gov/articles/PMC9088144/>
5. Shaffer F, Meehan ZM. *A Practical Guide to Resonance Frequency Assessment for HRV Biofeedback.* Frontiers in Neuroscience, 2020. <https://pmc.ncbi.nlm.nih.gov/articles/PMC7578229/>
6. Fincham GW, et al. *Effect of coherent breathing on mental health and wellbeing: a randomised placebo-controlled trial.* Scientific Reports, 2023. <https://www.nature.com/articles/s41598-023-49279-8>
7. Vierra J, et al. *Effects of sleep deprivation and 4-7-8 breathing control on heart rate variability, blood pressure, blood glucose, and endothelial function in healthy young adults.* Physiological Reports, 2022. <https://pmc.ncbi.nlm.nih.gov/articles/PMC9277512/>
8. Komori T. *The relaxation effect of prolonged expiratory breathing.* Mental Illness, 2018. <https://pmc.ncbi.nlm.nih.gov/articles/PMC6037091/>
9. Magnon V, Dutheil F, Vallet GT. *Benefits from one session of deep and slow breathing on vagal tone and anxiety in young and older adults.* Scientific Reports, 2021. <https://www.nature.com/articles/s41598-021-98736-9>
10. *Do Longer Exhalations Increase HRV During Slow-Paced Breathing?* Applied Psychophysiology and Biofeedback, 2024. <https://pmc.ncbi.nlm.nih.gov/articles/PMC11310264/>
11. Zhang X, et al. *The Effect of Slow-Paced Breathing on Cardiovascular and Emotion Functions: A Meta-Analysis and Systematic Review.* Mindfulness, 2023. <https://link.springer.com/article/10.1007/s12671-023-02294-2>
12. Ma X, et al. *The Effect of Diaphragmatic Breathing on Attention, Negative Affect and Stress in Healthy Adults.* Frontiers in Psychology, 2017. <https://www.frontiersin.org/journals/psychology/articles/10.3389/fpsyg.2017.00874/full>

*`calm` is a toy, not a medical device. Nothing here is medical advice.*

## Make targets

| target      | does |
|-------------|------|
| `build`     | build `./calm` (stripped binary) |
| `run`       | build and run |
| `install`   | build and install to `$BINDIR` (default `~/.local/bin`) |
| `uninstall` | remove the installed binary |
| `vet`       | `go vet ./...` |
| `clean`     | remove `./calm` |

## License

MIT — see [LICENSE](LICENSE).
