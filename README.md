# Crowd-GPT

Crowd-GPT explores how the Wisdom of Crowds approach, a principle that suggests a diverse group of individuals can often make better collective decisions than individual experts, can be applied to improve the intelligence of Generative Pretrained Transformer (GPT) models.

GPT models, powered by machine learning algorithms, are designed to understand and generate human-like text based on the input they receive. However, by introducing a 'crowd' of diverse personas akin to a human crowd, we can tremendously enhance their ability to generate diversified and more encompassing responses, thereby significantly improving their overall intelligence and prediction abilities.

## Key Components:

### Diverse Personas:
Just like humans, GPT will also use inputs from a variety of personas, bringing distinct perspectives. This includes 'experts' with thorough knowledge, 'generalists' providing broad viewpoints, 'outsiders' offering fresh insights, 'contrarians' presenting challenging viewpoints, 'comparable industry experts' lending a similar but uniquely nuanced perspective, and 'everyday users' representing common or popular points of view.

### Independent and Decentralized Input:
Each 'persona' contributes independently, avoiding common biases that may sway the group's opinion in actual crowd scenarios. This allows for each entry to be genuinely distinctive, adding to the overall diversity of the final output.

While there is great potential in utilizing the Wisdom of Crowds approach for enhancing GPT's capabilities, it's essential to consider the potential for biases and misinformation in inputs to be minimized and counter-checked.

This repo hosts tests, experiments, and develops models to leverage collective intelligence for GPT. We encourage contributions, suggestions, and collaborations to tap into the full potential of 'crowd wisdom' in GPT.

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.



## Approach

1.  Create 3 persons each of the six persona types for the Wisdom of the Crowds
 - Opinions Specialists: Individuals who have professional or expert knowledge in a particular subject matter and therefore, can provide skilled insights.

 - Generalists: Individuals with broad, general knowledge who can offer multiple perspectives on a given topic.

 -  Outsiders: Individuals who are not intimately familiar with the topic, lending an unbiased perspective which can offer fresh or unconventional insights.

 -  Contrarians: Those who often hold opposing viewpoints from the majority, offering a challenging perspective that encourages the exploration of all facets of an issue.

 -  Comparable Industry Perspective: Those who are experienced or knowledgeable in comparable industries can add a different viewpoint to the common knowledge of the crowd

 -  Everyday Users: Depending on the topic, average individuals or consumers who will give an on-the-ground perspective. 

 2. Have each person answer the question and store their results.

 3. Have each person rank all of the answers into their top 3 choices.

 4. If there is a majority answer, output that answer. If there isn't a majority, have each person rank their top 2 from the top 3 and then respond with the majority answer.